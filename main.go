package main

import (
	"fmt"
	"time"
)

const (
	StringForNow = "2006-01-02 15:04:05"
)

func addData(dataChan chan int) {
	for i := 0; i < 1000; i++ {
		// fmt.Printf("addData 第 %v 次\n", i+1)
		dataChan <- i
	}
	fmt.Println("關閉 dataChan")
	close(dataChan)
}

func handleData(dataChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		v, ok := <-dataChan
		if !ok {
			// fmt.Println("一條算完了")
			break
		}
		sum := 0
		for i := 1; i <= v; i++ {
			sum += i
		}
		var value = map[int]int{v: sum}
		resChan <- value
		// fmt.Println("handle 一次", v)
	}
	exitChan <- true
}

func main() {
	// 在產生資料時處理資料也在同步取出資料的話，上限數目就沒必要開到總數那麼大了
	var dataChan = make(chan int, 10)
	var resChan = make(chan map[int]int, 1000)
	var exitChan = make(chan bool, 4)

	start := time.Now().UnixMicro()
	// 開一條線產生基本資料
	go addData(dataChan)
	for i := 0; i < 4; i++ {
		// 開四條線做資料處理，所以 exitChan 記錄四個值當作開關使用
		go handleData(dataChan, resChan, exitChan)
	}

	// 在匿名函數最後使用 () 直接呼叫使用
	func() {
		// 主線程中控制結束時間的開關
		// 從 exitChan 取得四次值後繼續後面的邏輯
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	// 處理 resChan
	result := make(map[int]int, 1000)
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		for i, v := range v {
			result[i] = v
		}
	}
	//
	end := time.Now().UnixMicro()
	fmt.Printf("time = %v\n", end-start)
	for i := 0; i < 1000; i++ {
		fmt.Printf("result[%v] = %v\n", i, result[i])
	}

	fmt.Println("main end")
}
