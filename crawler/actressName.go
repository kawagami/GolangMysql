package crawler

import (
	"fmt"
	"regexp"
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
	// regex
	pattern := `^[a-zA-Z]{2,6}-[0-9]{2,6}`
	re, _ := regexp.Compile(pattern)
	//
	for _, video := range pathSlice {
		regexTitle := re.FindString(video.Title)
		if !va.Exist(regexTitle) && regexTitle != "" {
			// 取得名字
			var insert = mysql.VideoActresses{Title: regexTitle}
			if actressName := GetActressName(regexTitle); actressName != "" {
				// 建立要插入 DB的資料
				insert.Actress = actressName
				// 將資料寫入 DB
			} else {
				insert.Actress = ""
				fmt.Println("查無 actress name")
			}
			va.Insert(insert)
			fmt.Println("寫入", regexTitle)
			// 避免過度 request 被擋
			fmt.Println("待機 5 秒")
			time.Sleep(time.Second * 5)
		} else {
			fmt.Println("DB 有", regexTitle, "的資料")
		}
	}
}
