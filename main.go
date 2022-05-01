package main

import (
	"fmt"
	vtu "mods/videotidyup"
	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	res := vtu.GetVideos()
	for _, v := range res {
		fmt.Printf("檔名 = %v\n路徑 = %v\n\n", v.Name, v.Path)
	}
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
