package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	baseUrl = "https://javdb.com"
)

/*
依照 input 的番號
return 在 baseUrl 找到的第一個目標路徑
*/
func GetVideoInfo(videoName string) (result []string) {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	// c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
	// 	fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	// })

	c.OnHTML("a[href].box", func(e *colly.HTMLElement) {
		if e.Attr("href") != "" {
			firstPath := baseUrl + e.Attr("href")
			result = append(result, firstPath)
		}
		// fmt.Println(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	url := fmt.Sprintf("%v/search?q=%v&f=all", baseUrl, videoName)
	c.Visit(url) // Visit 要放最後

	return
}

// func GetVideoInnerInfo(path string) (result InfoWebRawData) {
func GetVideoInnerInfo(path string) (result InfoWebRawData) {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	// c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
	// 	fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	// })

	// 大標題
	// c.OnHTML("h2.title", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// 內頁的一欄一欄資訊集合體
	c.OnHTML("span.value", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		result.InfoStrings = append(result.InfoStrings, e.Text)
	})

	// 封面圖
	c.OnHTML(".video-cover", func(e *colly.HTMLElement) {
		// fmt.Println(e.Attr("src"))
		// fmt.Println(e.Text)
		result.Cover = e.Attr("src")
		// result.RawHtml = e
	})

	// 原始資料
	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
		result.RawHtml = string(r.Body)
	})

	// 包含日文字的標題
	c.OnHTML("h2.title", func(e *colly.HTMLElement) {
		// fmt.Println(e.Attr("src"))
		// fmt.Println(e.Text)
		result.LongTitle = strings.TrimSpace(e.Text)
		// result.RawHtml = e
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(path) // Visit 要放最後

	return
}

func GetActressName(fileTitle string) (actressName string) {
	paths := GetVideoInfo(fileTitle)
	if len(paths) > 1 {
		// 爬第二層
		rawData := GetVideoInnerInfo(paths[0])
		// 取得沒特定標籤可鎖定的資料並整理
		if len(rawData.InfoStrings) > 5 {
			infoIndex := len(rawData.InfoStrings) - 1
			res := strings.Split(rawData.InfoStrings[infoIndex], "♀")
			// 清除左右的空白
			aName := strings.TrimSpace(res[0])
			actressName = aName
			// fmt.Println(aName)
		} else {
			fmt.Println("取不到內頁資料", fileTitle)
		}
	} else {
		fmt.Println("第一層無資料")
	}
	return
}

func GetNews(url string) {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	// c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
	// 	fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	// })

	c.OnHTML("._1Zdp", func(e *colly.HTMLElement) {
		time := e.ChildText("time")
		location := e.ChildText(".theme-sub-cat")
		context := e.ChildText("h3")
		fmt.Printf("時間 = %v 地區 = %v 內容 = %v\n", time, location, context)
		// fmt.Println("")
		// fmt.Println(e.DOM.Find())
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	// url := fmt.Sprintf("%v/search?q=%v&f=all", baseUrl, videoName)

	c.Visit(url) // Visit 要放最後
}
