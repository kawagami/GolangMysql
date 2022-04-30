package zipdir

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 壓縮檔案
func ZipSource(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

// 放資料進入 channel
func PutDataIntoChannel(files map[string]string, newChan chan map[string]string) {
	for i, v := range files {
		newChan <- map[string]string{i: v}
	}
	close(newChan)
}

// 將資料從 channel 取出並處理
func Handle(newChan chan map[string]string, exitChan chan bool) {
	for {
		v, ok := <-newChan
		if !ok {
			break
		}
		for fileName, filePath := range v {
			// 檢查在 root 是否有同名壓縮檔
			sameName := filePath + ".zip"
			if _, err := os.Stat(sameName); os.IsNotExist(err) {
				// 沒同名壓縮檔的處理
				// 檢查資料夾下是圖片還是壓縮檔
				zipFileName := filePath + ".zip"
				// 這邊先偷懶，檔案數不等於 1 的都當是圖片檔
				if zfn, isZipFile := isZipFile(filePath); isZipFile {
					oldPath := filePath + `\` + zfn
					newPath := filePath + `.zip`
					fmt.Printf("move %v\n", fileName)
					os.Rename(oldPath, newPath)
				} else {
					// 壓縮檔案的處理
					fmt.Printf("zip %v\n", fileName)
					ZipSource(filePath, zipFileName)
				}
				// 壓縮完後移除資料夾
				checkExist(zipFileName, filePath)
			} else {
				// 有同名壓縮檔的處理
				fmt.Printf("有同名的壓縮檔 %v\n", fileName)
			}
		}
	}
	fmt.Println("one goroutine done!")
	exitChan <- true
}

// 檢查是不是壓縮檔
func isZipFile(path string) (string, bool) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	res1 := info[0].Name()
	res2 := len(info) == 1
	return res1, res2
}

// 將資料夾轉換成壓縮檔後使用，判斷該資料夾檔名的壓縮檔是否存在並處理
func checkExist(filePath, dirPath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 沒該檔案的動作
	} else {
		// 有該檔案的動作
		err := os.RemoveAll(dirPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("delete %v\n", dirPath)
	}
}

// const (
// 	root           = `D:\temp`
// 	stringForNow   = "2006-01-02 15:04:05"
// 	handleChannels = 10
// )

// func main() {
// 	fmt.Println(time.Now().Format(stringForNow))
// 	// 取得要整理的子資料夾 map
// 	files := zip.GetDirPath(root)
// 	//
// 	timeRecordStart := time.Now().UnixMicro()
// 	// 將資料放入 channel
// 	var newChan = make(chan map[string]string, 100)
// 	go zip.PutDataIntoChannel(files, newChan)
// 	// 把 channel 中的資料取出並處理
// 	var exitChan = make(chan bool, handleChannels)
// 	for i := 0; i < handleChannels; i++ {
// 		go zip.Handle(newChan, exitChan)
// 	}
// 	// 等待開啟的 goroutine 都結束
// 	for i := 0; i < handleChannels; i++ {
// 		<-exitChan
// 	}
// 	timeRecordEnd := time.Now().UnixMicro()
// 	timeRecordMicro := timeRecordEnd - timeRecordStart
// 	timeRecordSecond := float64(timeRecordMicro) / 1000000
// 	fmt.Printf("經過 %v 秒\n", timeRecordSecond)
// 	//
// 	fmt.Println("main end")
// }
