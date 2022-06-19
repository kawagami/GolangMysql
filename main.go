package main

import (
	"fmt"
	"mods/tidy"
	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	path := `D:\temp`
	res := tidy.GetAuthorName(path)
	// 紀錄所需的 map
	var dataLen = len(res)
	var data = make(map[string]string, dataLen)
	//
	for _, tidy := range res {
		data[tidy.Title] = tidy.Path
	}
	//
	for shouldBeAuthor, dataPath := range data {
		fmt.Println(shouldBeAuthor)
		fmt.Println(dataPath)
		fmt.Println("")
	}
	//
	fmt.Println("data len =", len(data))
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
