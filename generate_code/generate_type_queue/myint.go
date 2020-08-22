//go:generate ./generated_types/queue MyInt
package main

import "fmt"

// example of how to generate a code - struct and its methods
// cd generated_types
// go build queue.go
// cd ..
// go generate
// go run myint.go myint_queue.go

type MyInt int

func main() {
	var one, two, three MyInt = 1, 2, 3
	q := NewMyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Printf("First value: %d\n", q.Remove())
}
