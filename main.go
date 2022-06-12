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
	tidy.MoveFileToDest()
	//
	// res := tidy.ComicAuthors()
	// for _, t := range res {
	// 	fmt.Println(t.Title)
	// 	fmt.Println(t.Path)
	// 	fmt.Println("")
	// }
	//
	// temps := tidy.TempComics()
	// for _, temp := range temps {
	// 	fmt.Println(temp.Title)
	// 	fmt.Println(temp.Path)
	// 	fmt.Println("")
	// }
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
