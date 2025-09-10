package main

import (
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}
	defer conn.Close()
}
