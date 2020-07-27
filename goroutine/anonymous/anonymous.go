package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("outside goroutine")
	go func() {
		fmt.Println("inside goroutine")
	}()
	runtime.Gosched()
	fmt.Println("outside goroutine")
}
