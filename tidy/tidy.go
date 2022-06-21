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
	var needToFilter = []string{"汉化", "漢化", "翻訳"}
	//
	for _, file := range files {
		title := file.Name()
		dPath := path + `\` + file.Name()
		// 有匹配的話
		if find := re.FindAllStringSubmatch(title, -1); len(find) > 0 {
			// 對 [][]string 做 for range
			for _, ts := range find {
				// 過濾包含特定關鍵字的資料
				if IsStringSliceIncludeString(needToFilter, ts[1]) {
					continue
				}
				//
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

/*
檢查 []string 中的每個 string 是否存在於 inputString
*/
func IsStringSliceIncludeString(stringSlice []string, inputString string) bool {
	for _, ss := range stringSlice {
		// 檢查 inputString (regexp 抓出來的檔案名稱) 是否包含要過濾的字詞
		if strings.Contains(inputString, ss) {
			return true
		}
	}
	return false
}
