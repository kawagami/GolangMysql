package main

import (
	"fmt"
	zip "mods/zipDir"
	"time"
)

const (
	stringForNow = "2006-01-02 15:04:05"
	root         = `D:\temp`
)

func putDataIntoChannel(files map[string]string, newChan chan map[string]string) {
	for i, v := range files {
		newChan <- map[string]string{i: v}
	}
	close(newChan)
}

func handle(newChan chan map[string]string, exitChan chan bool) {
	for {
		v, ok := <-newChan
		if !ok {
			break
		}
		for fileName, filePath := range v {
			fmt.Printf("取出 %v 了\nPath = %v\n\n", fileName, filePath)
		}
		time.Sleep(time.Millisecond * 100)
	}
	exitChan <- true
}

func main() {
	fmt.Println(time.Now().Format(stringForNow))
	//
	files := zip.GetDirPath(root)
	//
	var newChan = make(chan map[string]string, 100)
	go putDataIntoChannel(files, newChan)
	var exitChan = make(chan bool, 4)
	for i := 0; i < 4; i++ {
		go handle(newChan, exitChan)
	}
	//
	for i := 0; i < 4; i++ {
		<-exitChan
	}
	//
	fmt.Println("main end")
}
