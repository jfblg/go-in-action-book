package main

import "fmt"

func main() {
	defer goodbey()
	fmt.Println("Hello world")
}

func goodbey() {
	fmt.Println("Goodbye")
}
