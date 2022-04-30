package main

import (
	"fmt"
	zip "mods/zipDir"
	"time"
)

const (
	root           = `D:\temp`
	stringForNow   = "2006-01-02 15:04:05"
	handleChannels = 10
)

func main() {
	fmt.Println(time.Now().Format(stringForNow))
	// 取得要整理的子資料夾 map
	files := zip.GetDirPath(root)
	//
	timeRecordStart := time.Now().UnixMicro()
	// 將資料放入 channel
	var newChan = make(chan map[string]string, 100)
	go zip.PutDataIntoChannel(files, newChan)
	// 把 channel 中的資料取出並處理
	var exitChan = make(chan bool, handleChannels)
	for i := 0; i < handleChannels; i++ {
		go zip.Handle(newChan, exitChan)
	}
	// 等待開啟的 goroutine 都結束
	for i := 0; i < handleChannels; i++ {
		<-exitChan
	}
	timeRecordEnd := time.Now().UnixMicro()
	timeRecordMicro := timeRecordEnd - timeRecordStart
	timeRecordSecond := float64(timeRecordMicro) / 1000000
	fmt.Printf("經過 %v 秒\n", timeRecordSecond)
	//
	fmt.Println("main end")
}
