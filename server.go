package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	id   int
	conn net.Conn
	ch   chan string
}

var (
	clients    = make(map[int]*Client)
	mu         sync.Mutex
	nextClient = 1
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// Register client
	mu.Lock()
	id := nextClient
	nextClient++
	client := &Client{ id: id, conn: conn, ch: make(chan string, 16) }
	clients[id] = client
	mu.Unlock()

	// Start writer goroutine
	go clientWriter(client)

	// Broadcast join (other clients only)
	broadcast(fmt.Sprintf("User [%d] joined", id), id)
	fmt.Println("User connected:", id)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		msg := fmt.Sprintf("User [%d]: %s", id, text)
		broadcast(msg, id)
		fmt.Println("recv:", msg)
	}

	// Cleanup
	conn.Close()
	mu.Lock()
	delete(clients, id)
	mu.Unlock()

	broadcast(fmt.Sprintf("User [%d] left", id), id)
	fmt.Println("User disconnected:", id)
}

func broadcast(msg string, except int) {
	mu.Lock()
	defer mu.Unlock()

	for id, c := range clients {
		if id == except {
			continue
		}
		select {
		case c.ch <- msg:
		default:
		}
	}
}

func clientWriter(c *Client) {
	w := bufio.NewWriter(c.conn)
	for msg := range c.ch {
		_, err := w.WriteString(msg + "\n")
		if err != nil {
			break
		}
		if err := w.Flush(); err != nil {
			break
		}
	}
	c.conn.Close()
}
