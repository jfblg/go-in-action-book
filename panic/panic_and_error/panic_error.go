package main

import (
	"errors"
	"fmt"
)

// ErrDevideByZero used when a division by zero is detected
var ErrDevideByZero = errors.New("Devide by zero is not allowed")

func main() {
	res, err := preCheckDevide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Result: ", res)

	res = devide(2, 0)
	fmt.Println("Result: ", res)

}

func preCheckDevide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDevideByZero
	}
	return devide(a, b), nil

}

func devide(a, b int) int {
	return a / b
}
