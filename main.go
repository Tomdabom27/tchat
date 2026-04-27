package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	username := promptUsername()

	switch args[0] {
	case "host":
		port := "8080"
		if len(args) >= 2 {
			port = args[1]
		}
		StartServer(port, username)

	case "join":
		if len(args) < 2 {
			fmt.Println("Usage: join <ip:port>")
			os.Exit(1)
		}
		StartClient(args[1], username)

	default:
		printUsage()
		os.Exit(1)
	}
}

func promptUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "anonymous"
	}
	return name
}

func printUsage() {
	fmt.Println("tchat — terminal chat between two machines")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  tchat host [port]       Start as host (default port: 8080)")
	fmt.Println("  tchat join <ip:port>    Join a host")
}
