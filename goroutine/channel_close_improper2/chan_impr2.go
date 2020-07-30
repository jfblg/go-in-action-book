package main

import (
	"time"
)

func main() {
	ch := make(chan bool)
	until := time.After(600 * time.Millisecond)

	go send(ch)

	for {
		select {
		case <-ch: // when the ch channel is closed a "false" value is received.
			println("Got message.")
		case <-until:
			println("Timeout")
		default:
			println("** HI ***")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch) // closed channel always return channles nil value
	println("Sent and closed")
}
