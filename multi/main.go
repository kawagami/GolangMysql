package main

import (
	"fmt"
	"time"
)

func say(word string) {
	for i := 0; i < 50; i++ {
		fmt.Println(word, i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	var word byte = 'A'
	for i := 0; i < 5; i++ {
		var input string = string(word)
		go say(input)
		word++
	}
	time.Sleep(time.Second * 5)
	fmt.Println("main end")
}
