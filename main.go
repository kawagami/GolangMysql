package main

import (
	"mods/multi"
)

const (
	StringForNow = "2006-01-02 15:04:05"
)

func main() {
	var data = make(chan int, 5)
	var flag = make(chan bool, 1)
	//
	go multi.WriteData(data)
	go multi.ReadData(data, flag)
	//
	for {
		if _, ok := <-flag; !ok {
			break
		}
	}
}
