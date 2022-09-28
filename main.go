package main

import (
	"fmt"
	"mods/crawler"

	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	// ------ 整理 temp 資料夾 to 漫畫區域 ------------------------------------
	// tidy.MoveFileToDest()
	// ------ 從網站上找影片的資料寫入資料庫 ------------------------------------
	// crawler.CrawlerActressNameGormVersion(`E:\video\H\unArranged`)
	// crawler.CrawlerActressNameGormVersion(`C:\waitToArrange`)

	crawler.CrawlerActressNameGormVersion(`D:\`)

	// ------ create new table ------------------------------------
	// db := sqlGorm.GetSailDb()
	// db.AutoMigrate(&sqlGorm.VideoMix{})
	// ------ simple ineffective looking for multiple directories ------------------------------------
	// dirs := tidy.GetFilesNameOfDir(`D:\video\H\`)
	// for _, dir := range dirs {
	// 	crawler.CrawlerActressNameGormVersion(dir.Path)
	// }
	// 資料庫寫入測試
	// db := sqlGorm.GetSailDb()
	// var model sqlGorm.VideoMix
	// model.Actress = "Actress"
	// model.CoverImg = "CoverImg"
	// model.LongTitle = "LongTitle"
	// model.RawHtml = "RawHtml"
	// db.Create(&model)
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
