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
	url := `https://news.cnyes.com/news/cat/headline`
	crawler.GetNews(url)
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
