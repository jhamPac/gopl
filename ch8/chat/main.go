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
type client chan<- string

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
				c <- msg
			}

		case c := <-entering:
			clients[c] = true

		case c := <-leaving:
			delete(clients, c)
			close(c)
		}
	}
}

func handleConn(conn net.Conn) {
	outgoingMsg := make(chan string)
	go clientWriter(conn, outgoingMsg)

	who := sillyname.Generate()
	outgoingMsg <- "client[ID] " + who
	messages <- who + " has entered the chat room"
	entering <- outgoingMsg

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- outgoingMsg
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
