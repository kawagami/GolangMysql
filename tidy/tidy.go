package tidy

import (
	"io/ioutil"
)

type Tidy struct {
	Title string
	Path  string
}

// 輸入路徑，取該層的檔案 & 資料夾，不取更深的資料
func GetFilesNameOfDir(path string) (res []Tidy) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		title := file.Name()
		path := path + file.Name()
		data := Tidy{Title: title, Path: path}
		res = append(res, data)
	}
	return
}
