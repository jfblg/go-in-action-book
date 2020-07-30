package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Failed to open a port 1000")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept a connection")
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s\n", err)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
		conn.Close()
	}
	repsponse(data, conn)

}

func repsponse(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("Intentional failure"))
}
