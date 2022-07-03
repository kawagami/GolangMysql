package main

import (
	"fmt"
	"mods/crawler"
	"mods/sqlGorm"
	"strings"
	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	// ------------------------------------
	video := `https://javdb.com/v/gG1by`
	result := crawler.GetVideoInnerInfo(video)
	// 未處理字串s
	// index := len(result.InfoStrings) - 1
	// fmt.Println(result.InfoStrings[index])
	// // 封面圖
	// fmt.Println(result.Cover)
	// // 原始檔案
	// fmt.Println(len(result.RawHtml))
	// ------------------------------------
	//
	// 取得 DB
	db := sqlGorm.GetDb()
	//
	var model sqlGorm.VideoMix
	// 取得名字
	if len(result.InfoStrings) > 5 {
		infoIndex := len(result.InfoStrings) - 1
		res := strings.Split(result.InfoStrings[infoIndex], "♀")
		// 清除左右的空白
		model.Actress = strings.TrimSpace(res[0])
	}
	//
	model.CoverImg = result.Cover
	model.RawHtml = result.RawHtml
	//
	db.Create(&model)
	// // Migrate the schema
	// db.AutoMigrate(&sqlGorm.VideoMix{})
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
