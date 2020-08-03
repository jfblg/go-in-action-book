package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	logger := log.New(logfile, "example ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Println("Regular message")
	logger.Fatalln("Fatal error")
	logger.Println("This message will NOT be logged")
}
