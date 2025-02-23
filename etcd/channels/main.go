package main

import (
	"fmt"
)

func forLoop10(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func forLoop20(ch chan int) {
	for i := 11; i <= 20; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	// random number would prints in
	go forLoop10(ch)
	go forLoop20(ch)
	
	for i := range ch {
		fmt.Println(i)
	}
}
