package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "a string to say")
var slovak bool

func init() {
	flag.BoolVar(&slovak, "slovak", false, "Use slovak language.")
	flag.BoolVar(&slovak, "s", false, "Use slovak language.")
}

func main() {
	flag.Parse()
	if slovak == true {
		fmt.Printf("Ahoj %s\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}
}
