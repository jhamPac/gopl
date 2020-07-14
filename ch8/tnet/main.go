package main

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
}

const (
	// Port is the port number that the server listens on
	Port = ":9000"
)

// Open connects to a TCP address
func Open(addr string) (*bufio.ReadWriter, error) {
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

// HandleFunc is a function that handles incoming commands
type HandleFunc func(*bufio.ReadWriter)

// Endpoint provides an endpoint to other processess to send data
type Endpoint struct {
	listener net.Listener
	handler  map[string]HandleFunc

	m sync.RWMutex
}

// NewEndpoint creates a new Endpoint
func NewEndpoint() *Endpoint {
	return &Endpoint{
		handler: map[string]HandleFunc{},
	}
}

// AddHandleFunc adds a new function for handling incoming data
func (e *Endpoint) AddHandleFunc(name string, f HandleFunc) {
	e.m.Lock()
	e.handler[name] = f
	e.m.Unlock()
}

// Listen starts listening on the endpoint
func (e *Endpoint) Listen() error {
	var err error
	e.listener, err = net.Listen("tcp", Port)
	if err != nil {
		return errors.Wrapf(err, "unable to listen on port %s\n", Port)
	}
	log.Println("listen on", e.listener.Addr().String())
	for {
		log.Println("accept a connection request")
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("failed accepting a connection request:", err)
			continue
		}
		log.Println("handle incoming message")
		go e.handleMessages(conn)
	}
}

func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		log.Print("receive command ")
		cmd, err := rw.ReadString('\n')

		switch {
		case err == io.EOF:
			log.Println("reached EOF - close this connection.\n ***")
			return

		case err != nil:
			log.Println("\nerror reading command. got: '"+cmd+"'\n", err)
			return
		}

		cmd = strings.Trim(cmd, "\n")
		log.Println(cmd + "'")

		e.m.RLock()
		handleCommand, ok := e.handler[cmd]
		e.m.RUnlock()
		if !ok {
			log.Printf("command %q is not registered\n", cmd)
			return
		}
		handleCommand(rw)
	}
}

func handleStrings(rw *bufio.ReadWriter) {
	log.Print("receive STRING message")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("cannot read from connection.\n", err)
	}
	s = strings.Trim(s, "\n ")
	log.Println(s)
	_, err = rw.WriteString("Thank you.\n")
	if err != nil {
		log.Println("cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("flush failed", err)
	}
}

func handleGob(rw *bufio.ReadWriter) {
	log.Print("receive GOB data")
	var data complexData

	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("error decoding GOB data:", err)
		return
	}

	log.Printf("outer complexData struct: \n%#v\n", data)
	log.Printf("inner complexData struct: \n%#v\n", data.C)
}
