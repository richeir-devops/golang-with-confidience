package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			ch <- 42
			wg.Done()
		}()
	}

	wg.Wait()
}
