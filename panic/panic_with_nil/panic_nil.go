package main

import "errors"

func main() {
	panic(errors.New("Can't execute the program"))
	panic("You did the mistake again")
}
