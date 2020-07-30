package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic trapped: %s (%T)\n", err, err)
		}
	}()

	panicFunc()

}

func panicFunc() {
	panic(errors.New("PANIC!"))
}
