package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "localhost:8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	// Reader goroutine: print what server sends
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		// if server closes connection, exit
		os.Exit(0)
	}()

	// Writer loop: read stdin and send to server
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		_, err := fmt.Fprintln(conn, in.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "send error:", err)
			return
		}
	}
}
