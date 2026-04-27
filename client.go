package main

import (
	"fmt"
	"net"
)

// StartClient dials the given address and enters the shared chat loop.
func StartClient(addr string, username string) {
	fmt.Printf("Connecting to %s...\n", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", addr, err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", conn.RemoteAddr())
	fmt.Println(divider())

	chatLoop(conn, username)
}
