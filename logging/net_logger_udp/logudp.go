package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "127.0.0.1:5145", timeout)
	if err != nil {
		panic("Failed to connect to localhost:5145")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example prefix ", f)

	logger.Println("Regular message")
	logger.Panicln("Panic message")
}
