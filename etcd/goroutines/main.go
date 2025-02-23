package main

import (
	"fmt"
	"sync"
)

func loop1(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Loop 1 - %d\n", i)
}

func getLoop1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go loop1(i, wg)
	}
}

func loop2(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Loop 2 - %d\n", i)
}

func getLoop2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i <= 15; i++ {
		wg.Add(1)
		go loop2(i, wg)
	}
}

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("goroutines started")
	wg.Add(2)
	go getLoop1(&wg)
	go getLoop2(&wg)
	wg.Wait()
	// time.Sleep(100 * time.Millisecond)
	fmt.Println("goroutines end")
}