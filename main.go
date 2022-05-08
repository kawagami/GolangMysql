package main

import (
	"fmt"
	"mods/mysql"
	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	var va mysql.VideoActresses
	data := va.Get()
	count := 0
	for _, fi := range data {
		count++
		fmt.Println(fi.Title)
		fmt.Println(fi.Actress)
		fmt.Println("")
	}
	fmt.Println(count)
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
