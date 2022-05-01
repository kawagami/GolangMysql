package main

import (
	"fmt"
	"time"

	"mods/GetInfo"
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
	for _, v := range pathSlice {
		fmt.Printf("title = %v\npath = %v\n\n", v.Title, v.Path)
	}
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
