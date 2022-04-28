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
