package main

import (
	"fmt"
	"mods/multiExam"
)

const (
	StringForNow = "2006-01-02 15:04:05"
)

func main() {
	var numChan = make(chan int, 2000)
	var resChan = make(chan map[int]int, 2000)
	var exitChan = make(chan bool, 1)
	go multiExam.CreateNum(numChan)
	for i := 0; i < 8; i++ {
		go multiExam.Calculate(numChan, resChan, exitChan)
	}
	for {
		if b, _ := <-exitChan; b {
			break
		}
	}
	res := multiExam.Sort(resChan, exitChan)
	for i, v := range res {
		fmt.Println("index =", i, "value =", v)
	}
}
