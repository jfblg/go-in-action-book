package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	// Example of a required flag
	Name   string `short:"n" long:"name" description:"A name to print" required:"true"`
	Slovak bool   `short:"s" long:"slovak" description:"Use slovak language"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}

	if opts.Slovak {
		fmt.Printf("Ahoj %s\n", opts.Name)
	} else {
		fmt.Printf("Hello %s\n", opts.Name)
	}

}
