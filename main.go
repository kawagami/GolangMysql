package main

import (
	"fmt"
	"time"

	"mods/GetInfo"
	"mods/mysql"

	"github.com/gocolly/colly"
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
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155") // Visit 要放最後

	//
	// for _, v := range pathSlice {
	// 	fmt.Printf("title = %v\npath = %v\n\n", v.Title, v.Path)
	// }
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
