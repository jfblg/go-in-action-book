package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5145")
	if err != nil {
		panic("Failed to connect to localhost:5145")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example prefix ", f)

	logger.Println("Regular message")
	logger.Panicln("Panic message")
}
