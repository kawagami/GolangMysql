package main

import (
	// "fmt"
	"fmt"
	"mods/multiExam"
)

const (
	StringForNow = "2006-01-02 15:04:05"
)

func main() {
	// 暫時儲存資料用的 channel 長度"目測"是影響不大
	var numChan = make(chan int, 100)
	// 處理過資料的儲存 channel ，如果開多個 goroutine 同時寫入的話要考慮關閉的時間點
	var resChan = make(chan map[int]int, 2000)
	// 用來記錄開啟的 goroutine 結束的 channel
	var exitChan = make(chan bool, 8)
	// 建立初始資料
	go multiExam.CreateNum(numChan)
	// 開啟八個 go 處理初始資料
	for i := 0; i < 8; i++ {
		go multiExam.Calculate(numChan, resChan, exitChan)
	}
	// 使用匿名函數包裝邏輯，取出特定數量的資料
	func() {
		// 沒取出設定數量的資料的話似乎就會停在這邊
		for i := 0; i < 8; i++ {
			// 不需要取出的資料，就丟去虛空
			<-exitChan
		}
		fmt.Println("取出 8 個了")
		// 取出設定數量的話就能把紀錄處理後資料的 channel 關閉
		close(resChan)
	}()
	// 整理 resChan 的資料
	// 整理這邊的邏輯應該可以使用 goroutine 處理，待學習
	var result = make(map[int]int, 2000)
	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		for i, v := range res {
			result[i] = v
		}
	}
	for i := 0; i < 2000; i++ {
		fmt.Printf("res[%v] = %v\n", i, result[i])
	}
	fmt.Println("main 結束")
}
