package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// generator to convert discrete values to a stream
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Printf("%v\n", num)
	}

	// generator to repeatedly call functions
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStr := make(chan interface{})
		go func() {
			defer close(valueStr)
			for {
				select {
				case <-done:
					return
				case valueStr <- fn():
				}
			}
		}()
		return valueStr
	}

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

}
