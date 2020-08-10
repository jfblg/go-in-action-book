package main

import "fmt"

func getName() string {
	return "there"
}

func main() {
	name := getName()
	fmt.Println("Hi ", name)
}
