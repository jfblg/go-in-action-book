package main

import (
	"testing"
)

func TestMyWriter(t *testing.T) {
	//var _ io.Writer = &MyWriter{} // this way you see at compile time that your
	// type does not implement io.Writer interface
}
