package main

import (
	"fmt"
	"mods/setting"
	zip "mods/zipDir"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(setting.TIMENOW))
	// 取得要整理的子資料夾 map
	files := zip.GetDirPath(setting.TEMPPATH)
	//
	timeRecordStart := time.Now().UnixMicro()
	// 將資料放入 channel
	var newChan = make(chan map[string]string, 100)
	go zip.PutDataIntoChannel(files, newChan)
	// 把 channel 中的資料取出並處理
	var exitChan = make(chan bool, setting.HANDLECHANNELS)
	for i := 0; i < setting.HANDLECHANNELS; i++ {
		go zip.Handle(newChan, exitChan)
	}
	// 等待開啟的 goroutine 都結束
	for i := 0; i < setting.HANDLECHANNELS; i++ {
		<-exitChan
	}
	timeRecordEnd := time.Now().UnixMicro()
	timeRecordMicro := timeRecordEnd - timeRecordStart
	timeRecordSecond := float64(timeRecordMicro) / 1000000
	fmt.Printf("經過 %v 秒\n", timeRecordSecond)
	//
	fmt.Println("main end")
}
