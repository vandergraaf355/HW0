package main

import (
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		return
	}
}
