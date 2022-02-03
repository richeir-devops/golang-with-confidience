package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time    time.Time
	severty string
	message string
}

var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	defer func() {
		fmt.Println("running defer...")
		close(logCh)
	}()

	wg.Add(1)

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	doneCh <- struct{}{}

	wg.Wait()
}
func logger() {
	for {
		select {
		case entry := <-logCh:
			if entry.time.Year() < time.Now().Year() {
				break
			} else {
				fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severty, entry.message)
			}
		case <-doneCh:
			wg.Done()
			break
		default:
			break
		}
	}
}
