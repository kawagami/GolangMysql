package sort

import (
	"fmt"
)

/*
將 input 的 []int slice 用 bubble sort 排序成從小到大
*/
func Bubble(slice []int) []int {
	fmt.Printf("length = %v\ncontent = %v\n", len(slice), slice)
	//
	if len(slice) > 1 {
		// 兩兩比較，所以次數是 n - 1
		for i := 0; i < len(slice)-1; i++ {
			// 每一 round 少一個要比較的數，所以 - i
			for j := 0; j < len(slice)-1-i; j++ {
				var tempNumber int
				if slice[j] > slice[j+1] {
					tempNumber = slice[j]
					slice[j] = slice[j+1]
					slice[j+1] = tempNumber
				}
			}
		}
	}
	//
	fmt.Printf("result length = %v\nresult content = %v\n", len(slice), slice)
	//
	return slice
}
