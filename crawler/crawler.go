package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

const (
	baseUrl = "https://javdb.com"
)

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
		fmt.Println(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	url := fmt.Sprintf("%v/search?q=%v&f=all", baseUrl, videoName)
	c.Visit(url) // Visit 要放最後

	return
}

func GetVideoInnerInfo(path string) (result []string) {
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	// c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
	// 	fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	// })

	// 大標題
	// c.OnHTML("h2.title", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnHTML("span.value", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		result = append(result, e.Text)
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(path) // Visit 要放最後

	return
}
