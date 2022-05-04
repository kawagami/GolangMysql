package crawler

import (
	"fmt"
	"time"

	"mods/GetInfo"
	"mods/mysql"
)

func CrawlerActressName(path string) {
	// call 使用 regexp 過濾檔案名稱的方法
	pathSlice := GetInfo.GetFileNumberFromDir(path)
	// 取得 DB 資料
	var va mysql.VideoActresses
	//
	for index, videoNumber := range pathSlice {
		fmt.Printf("進度 %v/%v\n", index+1, len(pathSlice))
		if !va.Exist(videoNumber) && videoNumber != "" {
			// 取得名字
			var insert = mysql.VideoActresses{Title: videoNumber}
			if actressName := GetActressName(videoNumber); actressName != "" {
				// 建立要插入 DB的資料
				insert.Actress = actressName
				// 將資料寫入 DB
			} else {
				insert.Actress = ""
				fmt.Println("查無 actress name")
			}
			fmt.Println("寫入", videoNumber)
			va.Insert(insert)
			// 最後一個就不待機了
			if index+1 < len(pathSlice) {
				// 避免過度 request 被擋
				fmt.Println("待機 5 秒")
				time.Sleep(time.Second * 5)
			}
		} else if videoNumber != "" {
			fmt.Println("DB 有", videoNumber, "的資料")
		} else {
			fmt.Println("不符合影片檔名格式", videoNumber)
		}
	}
}
