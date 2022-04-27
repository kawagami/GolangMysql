package main

import (
	"fmt"
	"math/rand"
	"sort"

	// "mods/sort"

	"mods/structs"
)

func main() {
	// var number int
	// fmt.Println("請輸入要取幾個 fi 數")
	// fmt.Scanln(&number)
	// var slice []int = []int{1, 2, 3, 4, 5}
	// maxIndex := len(slice) - 1
	// fmt.Println(slice[maxIndex])
	// var slice []int = []int{498, 1531, 89789, 15, 3}
	var ss structs.StuSlice
	for i := 0; i < 10; i++ {
		stu := structs.Student{
			Name:  fmt.Sprintf("學生--%v", rand.Intn(100)),
			Score: rand.Intn(100),
		}
		ss = append(ss, stu)
	}
	for _, v := range ss {
		fmt.Printf("Name = %v\tScore = %v\n", v.Name, v.Score)
	}
	fmt.Println("排序後")
	sort.Sort(ss)
	for _, v := range ss {
		fmt.Printf("Name = %v\tScore = %v\n", v.Name, v.Score)
	}
}
