package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1) // a buffered channel must be used
	// sender can send a message to a buffered channel and move on
	for i := 0; i < 10; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)

}

func worker(id int, lock chan bool) {
	fmt.Printf("%d is trying to lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing to lock\n", id)
	<-lock
}
