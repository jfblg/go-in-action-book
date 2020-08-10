package main

import (
	"fmt"
	"io"
)

type MyWriter struct {
	Name string
}

// It is intention that io.Writer interface is not implemented here
// to demonstrate how it can be detected in a test by a technique
// called canary test
func (m *MyWriter) Write([]byte) error {
	return nil
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}
	fmt.Println(m)
}

func doSomething(m map[string]interface{}) {
	w := m["w"].(io.Writer) // generates runtime exception
	fmt.Println(w)
}
