package main

import "log"

func main() {
	i := 10
	log.Println("Some log")      // to os.Stderr
	log.Printf("Value: %d\n", i) // to os.Stderr
	log.Fatalln("Fatal error")   // os.Stderr + os.Exit(1)
}
