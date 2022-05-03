package main

import (
	"fmt"
	"time"

	"mods/GetInfo"
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
	// record := va.Get()
	fmt.Println("檢查", va.Exist(pathSlice[0].Title))
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
