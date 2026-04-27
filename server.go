package main

import (
	"fmt"
	"net"
)

// StartServer binds a TCP listener on the given port, waits for exactly one
// client to connect, then enters the shared chat loop.
func StartServer(port string, username string) {
	addr := ":" + port
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Error starting server on %s: %v\n", addr, err)
		return
	}
	defer ln.Close()

	fmt.Printf("Hosting on port %s — waiting for someone to join...\n", port)
	fmt.Println("(Press Ctrl+C to quit)")

	conn, err := ln.Accept()
	if err != nil {
		fmt.Printf("Error accepting connection: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected: %s\n", conn.RemoteAddr())
	fmt.Println(divider())

	chatLoop(conn, username)
}
