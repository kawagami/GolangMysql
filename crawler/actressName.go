package crawler

import (
	"fmt"
	"time"

	"mods/GetInfo"
	"mods/mysql"
)

func CrawlerActressName() {
	path := `C:\waitToArrange`
	var pathSlice []GetInfo.VideoCrawler
	err := GetInfo.GetDir(path, &pathSlice)
	if err != nil {
		panic(err)
	}
	//
	// 取得 DB 資料
	var va mysql.VideoActresses
	for _, video := range pathSlice {
		if !va.Exist(video.Title) {
			// 取得名字
			var insert = mysql.VideoActresses{Title: video.Title}
			if actressName := GetActressName(video.Title); actressName != "" {
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
}
