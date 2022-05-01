package main

import (
	"fmt"
	"time"

	"mods/cmdsomething"
)

func main() {
	fmt.Println("")
	start := time.Now().UnixMicro()
	//
	cmdsomething.WebWithCopy()
	//
	end := time.Now().UnixMicro()
	timeResult := end - start
	processTime := float64(timeResult) / 1000000
	fmt.Printf("經過 %v 秒\n", processTime)
	fmt.Println("")

}
