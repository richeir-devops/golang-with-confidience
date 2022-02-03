package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
