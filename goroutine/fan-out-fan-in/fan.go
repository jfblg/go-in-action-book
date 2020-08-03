package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

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

	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		valueStreamInt := make(chan int)
		go func() {
			defer close(valueStreamInt)
			for v := range valueStream {
				select {
				case <-done:
					return
				case valueStreamInt <- v.(int):
				}
			}
		}()
		return valueStreamInt
	}

	fanIn := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		// wait for all reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Intn(400000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream) // todo implement primeFinder
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

}
