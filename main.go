package main

import (
	"fmt"
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
	// fmt.Println("檢查", va.Exist(pathSlice[0].Title))
	for _, video := range pathSlice {
		if !va.Exist(video.Title) {
			// 取得名字
			var insert = mysql.VideoActresses{Title: video.Title}
			if actressName := crawler.GetActressName(video.Title); actressName != "" {
				// 建立要插入 DB的資料
				insert.Actress = actressName
				// 將資料寫入 DB
			} else {
				insert.Actress = ""
				fmt.Println("查無 actress name")
			}
			va.Insert(insert)
			fmt.Println("寫入", video.Title)
			// 避免過度 request 被擋
			fmt.Println("待機 5 秒")
			time.Sleep(time.Second * 5)
		} else {
			fmt.Println("DB 有", video.Title, "的資料")
		}
	}
	// 檢查 DB 資料存在與否的邏輯等待思慮周全
	// 可能在 VideoActresses 增加使用 sql 檢查存在的方法
	// for _, v := range pathSlice {
	// 	for _, dbData := range record {
	// 		if dbData.Title == v.Title {
	// 			fmt.Println("在資料庫內", v.Title)
	// 		} else {
	// 			fmt.Println("no data", v.Title)
	// 		}
	// 	}
	// }
	// an := crawler.GetActressName(targetName)
	// fmt.Println(an)
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
