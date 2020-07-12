package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jhampac/gopl/ch8/ftpd"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 9000, "listen port")

	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("opening main listener:", err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Print("accepting new connection:", err)
		}
		go ftpd.NewConn(c).Run()
	}
}
