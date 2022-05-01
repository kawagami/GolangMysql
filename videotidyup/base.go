package videotidyup

import (
	"io/ioutil"
	"strings"
)

type SimVideo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

const (
	root = `D:\`
)

func GetVideos() (res []SimVideo) {
	fiSlice, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, v := range fiSlice {
		if strings.HasSuffix(v.Name(), ".mp4") {
			// fmt.Println(v.Name))
			path := root + v.Name()
			data := SimVideo{Name: v.Name(), Path: path}
			res = append(res, data)
		}
	}
	return
}
