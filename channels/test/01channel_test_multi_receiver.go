package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var ch1 = make(chan int, 10)

func main() {
	wg.Add(2)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
	}(ch1)

	go func(ch <-chan int) {
		fmt.Printf("Thread1: %v\n", <-ch)
		wg.Done()
	}(ch1)

	go func(ch <-chan int) {
		fmt.Printf("Thread2: %v\n", <-ch)
		wg.Done()
	}(ch1)

	wg.Wait()
}
