package main

import (
	"fmt"
	"strings"
	"time"

	"mods/GetInfo"
	"mods/crawler"
	"mods/mysql"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	path := `C:\waitToArrange`
	var pathSlice []GetInfo.VideoCrawler
	err := GetInfo.GetDir(path, &pathSlice)
	if err != nil {
		panic(err)
	}
	//
	// 取得 DB 資料
	var va mysql.VideoActresses
	record := va.Get()
	fmt.Println(len(record))
	//
	targetName := pathSlice[30].Title
	paths := crawler.GetVideoInfo(targetName)
	if len(paths) > 1 {
		// 爬第二層
		datas := crawler.GetVideoInnerInfo(paths[0])
		// 取得沒特定標籤可鎖定的資料並整理
		if len(datas) > 5 {
			infoIndex := len(datas) - 1
			res := strings.Split(datas[infoIndex], "♀")
			// 清除左右的空白
			aName := strings.TrimSpace(res[0])
			fmt.Println(aName)
		} else {
			fmt.Println("取不到內頁資料", targetName)
		}
	} else {
		fmt.Println("第一層無資料")
	}
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
