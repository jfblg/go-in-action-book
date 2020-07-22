package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func simpleExecution() {
	go count()
	time.Sleep(time.Millisecond)
	fmt.Println("Hello there")
	time.Sleep(time.Millisecond * 400)
}

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, "  ")
	}
}

func execPrintCount() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, -2, 10}
	go printCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of execPrintCOunt")
}

func main() {
	// simpleExecution()
	execPrintCount()
}
