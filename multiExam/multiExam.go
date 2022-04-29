package multiExam

import "fmt"

func CreateNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	fmt.Println("CreateNum 結束")
	close(numChan)
}

func Calculate(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		value, ok := <-numChan
		if !ok {
			break
		}
		mapValue := 0
		for i := 0; i <= value; i++ {
			mapValue += i
		}
		// fmt.Printf("key = %v value = %v\n", value, mapValue)
		resValue := map[int]int{value: mapValue}
		resChan <- resValue
	}
	fmt.Println("一個 goroutine 結束")
	exitChan <- true
}

func Sort(resChan chan map[int]int, exitChan chan bool) (res [2000]int) {
	for i, v := range <-resChan {
		fmt.Printf("key = %v value = %v\n", i, v)
		res[i] = v
	}
	return
}
