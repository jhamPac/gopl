package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

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

const timeout = 5 * time.Second

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
	out := make(chan string)
	go clientWriter(conn, out)
	in := make(chan string)
	go clientReader(conn, in)

	var who string
	nameTimer := time.NewTimer(15 * time.Second)
	out <- "Enter your name"

	select {
	case name := <-in:
		who = name
	case <-nameTimer.C:
		who = sillyname.Generate()
	}

	c := client{out, who}
	out <- "chat name: " + who
	messages <- who + " has entered the chat room"
	entering <- c
	idle := time.NewTimer(timeout)

Loop:
	for {
		select {
		case msg := <-in:
			messages <- who + ": " + msg
			idle.Reset(timeout)

		case <-idle.C:
			conn.Close()
			break Loop
		}
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

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}
