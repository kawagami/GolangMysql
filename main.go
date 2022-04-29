package main

import (
	// "fmt"
	"fmt"
	"time"
)

const (
	StringForNow = "2006-01-02 15:04:05"
)

func addData(dataChan chan int) {
	for i := 0; i < 50; i++ {
		fmt.Printf("addData 第 %v 次\n", i+1)
		dataChan <- i
	}
	fmt.Println("關閉 dataChan")
	close(dataChan)
}

func main() {
	var dataChan = make(chan int, 10)

	go addData(dataChan)

	func() {
		for {
			v, ok := <-dataChan
			if !ok {
				fmt.Println("沒東西了")
				break
			}
			fmt.Println("取得", v)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	fmt.Println("main end")
}
