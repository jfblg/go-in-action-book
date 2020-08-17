package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := 1
	if isStringer(a) {
		fmt.Printf("%T is a stringer\n", a)
	} else {
		fmt.Printf("%T is a NOT stringer\n", a)
	}

	b := bytes.NewBuffer([]byte("Hi there"))
	if isStringer(b) {
		fmt.Printf("%T is a stringer\n", b)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer) // type assertion to desired interface
	return ok

}
