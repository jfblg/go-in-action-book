package main

import (
	"math/rand"
	"runtime"
	"time"
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

	fanIn := func(
		done <-chan interface,
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func( c <-chan interface{}) {
			defer wg.Done()
			for i:= range c {
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

	start := time.Now()
	rand := func() interface{} { return rand.Intn(400000) }
	randIntStream := toInt(done, repeatFn(done, rand))
	// todo toInt stage

	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	

}
