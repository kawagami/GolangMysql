package Fibonacci

import (
	"fmt"
)

func Fibonacci() (slice []int) {
	var number int
	fmt.Println("請輸入要取幾個 fi 數")
	fmt.Scanln(&number)
	slice = []int{1, 1}
	for i := 1; i < number-1; i++ {
		var index int = i + 1
		slice = append(slice, slice[index-2]+slice[index-1])
	}
	for i, v := range slice {
		index := i + 1
		fmt.Printf("第 %v 個是 %v\n", index, v)
	}
	//
	return
}
