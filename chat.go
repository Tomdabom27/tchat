package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// chatLoop is the symmetric chat session shared by both host and client.
// It spawns two goroutines:
//   - one reads from stdin and writes to the connection
//   - one reads from the connection and prints to stdout
//
// Either side typing "/quit" or closing the terminal ends the session.
func chatLoop(conn net.Conn, username string) {
	done := make(chan struct{})

	// Goroutine 1: remote → stdout
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("\r%s\n> ", line) // overwrite the "> " prompt
		}
		fmt.Println("\n[Connection closed by remote]")
		close(done)
	}()

	// Goroutine 2: stdin → remote
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())

			if text == "/quit" {
				fmt.Println("[You left the chat]")
				conn.Close()
				return
			}

			if text == "" {
				fmt.Print("> ")
				continue
			}

			msg := formatMessage(username, text)
			fmt.Fprintln(conn, msg)
			fmt.Print("> ")
		}
		// stdin closed (Ctrl+D)
		conn.Close()
	}()

	<-done
}

// formatMessage produces a timestamped, attributed message line.
func formatMessage(username, text string) string {
	ts := time.Now().Format("15:04:05")
	return fmt.Sprintf("[%s] %s: %s", ts, username, text)
}

// divider returns a visual separator line.
func divider() string {
	return strings.Repeat("─", 40)
}
