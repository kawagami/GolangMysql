package GetInfo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type VideoCrawler struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

/*
遍歷所有層的資料夾和檔案
*/
func GetAllFile(files *[]string) filepath.WalkFunc {
	/* 範例
	var files []string
	root := "D:\\temp"
	err := filepath.Walk(root, GetInfo.GetAllFile(&files))
	if err != nil {
		panic(err)
	}
	*/
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("visit 的 path = %v\n", path)
		*files = append(*files, path)
		return nil
	}
}

/*
只遍歷下一層的所有資料夾
*/
func GetDir(pathname string, pathSlice *[]string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		*pathSlice = append(*pathSlice, fi.Name())
	}
	return err
}

func GetFileNumberFromDir(path string) (response []string) {
	var pathSlice []string
	err := GetDir(path, &pathSlice)
	if err != nil {
		panic(err)
	}
	pattern := `^([a-zA-Z]{2,6}-[0-9]{2,6}).*(\..*)$`
	re := regexp.MustCompile(pattern)
	for _, fileName := range pathSlice {
		if result := re.FindStringSubmatch(fileName); len(result) > 1 {
			response = append(response, result[1])
			// fmt.Printf("原檔案名 = %v\n番號 = %v\n檔案類型 = %v\n\n", result[0], result[1], result[2])
		}
	}
	return
}
