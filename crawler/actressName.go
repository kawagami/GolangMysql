package crawler

import (
	"fmt"
	"strings"
	"time"

	"mods/GetInfo"
	"mods/mysql"
	"mods/sqlGorm"

	"gorm.io/gorm"
)

/*
透過 crawler 取得 input 路徑的影片資料
*/
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

/*
使用 gorm 改寫上面的自製 function
path 是要查詢的資料夾路徑
*/
func CrawlerActressNameGormVersion(path string) {
	// call 使用 regexp 過濾檔案名稱的方法
	pathSlice := GetInfo.GetFileNumberFromDir(path)
	// // 取得 DB 連線
	// db := sqlGorm.GetDb()
	// 取得 sail 用的 DB 連線
	db := sqlGorm.GetSailDb()
	// defer db.Close()
	//
	for index, videoNumber := range pathSlice {
		fmt.Printf("進度 %v/%v\n", index+1, len(pathSlice))
		// 檢查資料庫是否已經存在
		if !IsTitleExists(videoNumber, db) && videoNumber != "" {
			// 爬蟲後有資料的話
			if data := crawleredData(videoNumber); data.Title != "" {
				// 將資料寫入 DB
				fmt.Println("寫入", videoNumber)
				db.Create(&data)
			} else {
				fmt.Println("查無", videoNumber, "的 actress name")
			}
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

/*
檢查 title 在資料庫內是否已經存在
*/
func IsTitleExists(title string, db *gorm.DB) (result bool) {
	// // 取得 DB 連線
	// db := sqlGorm.GetDb()
	// 取得 sail 用的 DB 連線
	// db := sqlGorm.GetSailDb()
	var model sqlGorm.VideoMix
	var exists bool
	db.Model(model).
		Select("count(*) > 0").
		Where("title = ?", title).
		Find(&exists)
	//
	return exists
}

func crawleredData(number string) (model sqlGorm.VideoMix) {
	// 番號
	path := GetVideoInfo(number)
	// 沒資料就停下
	if !(len(path) > 0) {
		return
	}
	// 有一筆以上的資料
	// video := `https://javdb.com/v/gG1by`
	video := path[0]
	result := GetVideoInnerInfo(video)
	// // 未處理字串s
	// index := len(result.InfoStrings) - 1
	// fmt.Println(result.InfoStrings[index])
	// // 封面圖
	// fmt.Println(result.Cover)
	// // 原始檔案
	// fmt.Println(len(result.RawHtml))
	//
	// 要存入資料庫的 model
	// 取得名字
	if len(result.InfoStrings) > 5 {
		infoIndex := len(result.InfoStrings) - 1
		res := strings.Split(result.InfoStrings[infoIndex], "♀")
		// 清除左右的空白
		// 演員名
		model.Actress = strings.TrimSpace(res[0])
	}
	// 取得 番號、封面圖、原始檔案
	// 番號
	model.Title = number
	// 封面圖
	model.CoverImg = result.Cover
	// 原始資料
	// model.RawHtml = result.RawHtml
	model.RawHtml = "strings too long, pause"
	// 含日文的長標題
	model.LongTitle = result.LongTitle
	//
	return
}
