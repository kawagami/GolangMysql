package multi

import (
	"fmt"
)

func WriteData(intChan chan int) {
	for i := 0; i < 10000; i++ {
		intChan <- i
		fmt.Println("write data", i)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", v)
	}
	exitChan <- true
	close(exitChan)
}

// func main() {
// 	var data = make(chan int, 5)
// 	var flag = make(chan bool, 1)
// 	//
// 	go multi.WriteData(data)
// 	go multi.ReadData(data, flag)
// 	//
// 	for {
// 		if _, ok := <-flag; !ok {
// 			break
// 		}
// 	}
// }
