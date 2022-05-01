package GetInfo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
func GetDir(pathname string, pathSlice *[]VideoCrawler) error {
	var tempSlice []VideoCrawler
	//
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		title := strings.TrimSuffix(fi.Name(), ".mp4")
		title = strings.TrimSuffix(title, ".wmv")
		title = strings.TrimSuffix(title, ".mkv")
		data := VideoCrawler{title, pathname + "\\" + fi.Name()}
		tempSlice = append(tempSlice, data)
		// fmt.Println("get", fi.Name())
	}
	*pathSlice = tempSlice
	return err
}
