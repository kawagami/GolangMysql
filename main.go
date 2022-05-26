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
	path := `E:\video\H\unArranged\`
	destPath := `E:\video\H\true\`
	tidy.TidyVideoByActressName(path, destPath, false)
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
