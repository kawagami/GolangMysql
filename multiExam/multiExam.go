package multiExam

import "fmt"

func CreateNum(numChan chan int) {
	for i := 0; i < 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func Calculate(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		value, ok := <-numChan
		if !ok {
			break
		}
		mapValue := 0
		for i := 1; i <= value; i++ {
			mapValue += i
		}
		fmt.Printf("key = %v value = %v\n", value, mapValue)
		resValue := map[int]int{value: mapValue}
		resChan <- resValue
	}
	if len(resChan) >= 2000 {
		exitChan <- true
		close(resChan)
		close(exitChan)
		fmt.Println("有進 IF")
	}
}

func Sort(resChan chan map[int]int, exitChan chan bool) (res [2000]int) {
	for i, v := range <-resChan {
		res[i] = v
	}
	return
}
