package main

import "fmt"

func main() {

	// stage example
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4, 5}

	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}

	multiplyStream := func(value, multiplier int) int {
		return value * multiplier
	}

	addStream := func(value, additive int) int {
		return value + additive
	}

	for _, v := range ints {
		fmt.Println(addStream(multiplyStream(v, 2), 1))
	}

	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()
		return intStream
	}

	multiplyCon := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- multiplier * i:
				}
			}
		}()
		return multipliedStream
	}

	addCon := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- additive + i:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiplyCon(done, addCon(done, multiplyCon(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

}
