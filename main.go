package main

import (
	"fmt"
	"mods/sqlGorm"
	"time"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	db := sqlGorm.GetDb()
	var vas []sqlGorm.VideoActress
	db.Find(&vas)
	for _, va := range vas {
		fmt.Println(va)
	}
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")
}
