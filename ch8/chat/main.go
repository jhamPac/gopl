package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/jhampac/gopl/ch8/chat/sillyname"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	log.Println("chat server listening on port: 9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// outgoing message channel
type client struct {
	Out  chan<- string // outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// broadcast message to all clients
			for c := range clients {
				c.Out <- msg
			}

		case c := <-entering:
			clients[c] = true
			c.Out <- "Present:"
			for c := range clients {
				c.Out <- c.Name
			}

		case c := <-leaving:
			delete(clients, c)
			close(c.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	oCh := make(chan string)
	go clientWriter(conn, oCh)

	who := sillyname.Generate()
	c := client{oCh, who}
	oCh <- "chat name: " + who
	messages <- who + " has entered the chat room"
	entering <- c

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- c
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
