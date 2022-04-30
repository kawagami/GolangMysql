package main

import (
	"fmt"
	"io/ioutil"
	zip "mods/zipDir"
	"os"
	"time"
)

const (
	// root           = `D:\temp`
	stringForNow   = "2006-01-02 15:04:05"
	root           = `C:\test`
	handleChannels = 10
)

func putDataIntoChannel(files map[string]string, newChan chan map[string]string) {
	for i, v := range files {
		newChan <- map[string]string{i: v}
	}
	close(newChan)
}

func handle(newChan chan map[string]string, exitChan chan bool) {
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
					zip.ZipSource(filePath, zipFileName)
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

func isZipFile(path string) (string, bool) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n\n", info[0].Name())
	// fmt.Printf("有 %v 個檔案", len(info))
	res1 := info[0].Name()
	res2 := len(info) == 1
	return res1, res2
}

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

func main() {
	fmt.Println(time.Now().Format(stringForNow))
	//
	files := zip.GetDirPath(root)
	//
	timeRecordStart := time.Now().UnixMicro()
	var newChan = make(chan map[string]string, 100)
	go putDataIntoChannel(files, newChan)
	var exitChan = make(chan bool, handleChannels)
	for i := 0; i < handleChannels; i++ {
		go handle(newChan, exitChan)
	}
	//
	for i := 0; i < handleChannels; i++ {
		<-exitChan
	}
	timeRecordEnd := time.Now().UnixMicro()
	//
	timeRecordMicro := timeRecordEnd - timeRecordStart
	timeRecordSecond := float64(timeRecordMicro) / 1000000
	fmt.Printf("經過 %v 秒\n", timeRecordSecond)
	//
	fmt.Println("main end")
}
