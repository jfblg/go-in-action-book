package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	// if statement can have optionall value assignment before the expression
	// common patter:
	if res, err := Concat(args...); err != nil { // args... - expands the array
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated arguments: '%s'\n", res)
	}

}

// Concat concatenantes a buch of strings
// It returns an empty string and an error if no strings were passed in
// p ...string - called varargs. Any numbe rof arguments of string type is accepted by the function.
// Go puts the in a slice of that type
func Concat(p ...string) (string, error) {
	if len(p) < 1 {
		return "", errors.New("No string arguments provided")
	}
	return strings.Join(p, " "), nil
}
