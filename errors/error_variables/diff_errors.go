package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// This example creates a package scoped variables
// which are assigne a random number, for distinguisching between various errors

// ErrTimeout used when connection timed out
var ErrTimeout = errors.New("The request timed out")

// ErrRejected used when connection was rejected
var ErrRejected = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(35))

func main() {
	response, err := SendRequest("Hi there")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = SendRequest("Hi")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

// SendRequest mocks a sender
func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
