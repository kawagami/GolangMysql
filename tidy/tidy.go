package tidy

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Tidy struct {
	Title string
	Path  string
}

type TidyArr struct {
	Titles [][]string
	Path   string
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

/*
在 temp 使用 regexp 取得未整理檔案名稱中的作者名稱
*/
func GetAuthorName(path string) (res []Tidy) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	//
	pattern := `(\[[^\[]*\])`
	re := regexp.MustCompile(pattern)
	//
	for _, file := range files {
		title := file.Name()
		dPath := path + `\` + file.Name()
		// 有匹配的話
		if find := re.FindAllStringSubmatch(title, -1); len(find) > 0 {
			// 對 [][]string 做 for range
			for _, ts := range find {
				// 過濾包含特定關鍵字的資料
				if strings.Contains(ts[1], "汉化") {
					continue
				}
				if strings.Contains(ts[1], "漢化") {
					continue
				}
				if strings.Contains(ts[1], "翻訳") {
					continue
				}
				var data Tidy
				data.Title = ts[1]
				data.Path = dPath
				res = append(res, data)
			}
		}
		//
	}
	return
}
