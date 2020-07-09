package ftpd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

// Conn is the main FTP interface
type Conn struct {
	rw           net.Conn
	dataHostPort string
	prevCmd      string
	pasvListener net.Listener
	cmdErr       error
	binary       bool
}

// NewConn creates and returns a pointer to conn
func NewConn(cmdConn net.Conn) *Conn {
	return &Conn{rw: cmdConn}
}

func hostPortToFTP(hostport string) (addr string, err error) {
	host, portStr, err := net.SplitHostPort(hostport)
	if err != nil {
		return "", err
	}
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "", err
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return "", err
	}
	ip := ipAddr.IP.To4()
	s := fmt.Sprintf("%d,%d,%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3], port/256, port%256)
	return s, nil
}

func hostPortFromFTP(address string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}

type logPairs map[string]interface{}

func (c *Conn) log(pairs logPairs) {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "addr=%s", c.rw.RemoteAddr().String())
	for k, v := range pairs {
		fmt.Fprintf(b, " %s=%s", k, v)
	}
	log.Print(b.String())
}

func (c *Conn) dataConn() (conn io.ReadWriteCloser, err error) {
	switch c.prevCmd {
	case "PORT":
		conn, err = net.Dial("tcp", c.dataHostPort)
		if err != nil {
			return nil, err
		}

	case "PASV":
		conn, err = c.pasvListener.Accept()
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("previous command not PASV or PORT")
	}
	return conn, nil
}
