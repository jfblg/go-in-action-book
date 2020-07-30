package main

// https://dev.to/williamhgough/how-to-test-tcpudp-connections-in-go---part-1-3bga

import "log"

type Server interface {
	Run() error
	Close() error
}

type TCPServer struct {
	addr string, 
	server net.Listener
}

func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil  {
		return
	}
	for {
		conn, err := t.server.Accept()
		if err != nil {
			err = errors.New("could not accept connection")
			break
		}

		if conn == nil {
			err = errors.New("could not create connection")
			break
		}
		conn.Close()
	}
	return
}

func (t *TCPServer) Close() error {
	retrun t.server.Close()
}


func NewServer(protocol, addr string) (Server, error) {
	switch strings.ToLower(protocol) {
	case "tcp":
		return &TCPServer{
			addr: addr
		}, nil
		case "udp":
	}
	return nil, errors.New("Invalid protocol given")
}

var srv Server

func init() {
	srv, err := NewServer("tcp", ":1234")
	if err != nil {
		log.Println("error starting TCP server")
		return
	}

	go func() {
		srv.Run()
	}()
}

func main() {

}
