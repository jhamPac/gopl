package ftpd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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

func (c *Conn) list(args []string) {
	var filename string
	switch len(args) {
	case 0:
		filename = "."
	case 1:
		filename = args[0]
	default:
		c.writeln("501 too many arguments")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		c.writeln("550 file not found")
		return
	}
	c.writeln("150 here comes the directory listing")
	w, err := c.dataConn()
	if err != nil {
		c.writeln("425 can't open data connection")
		return
	}
	defer w.Close()
	stat, err := file.Stat()
	if err != nil {
		c.log(logPairs{"cmd": "LIST", "err": err})
		c.writeln("450 requested file action not taken. file unavailable")
	}

	if stat.IsDir() {
		filenames, err := file.Readdirnames(0)
		if err != nil {
			c.writeln("550 can't read directory")
			return
		}
		for _, f := range filenames {
			_, err = fmt.Fprint(w, f, c.lineEnding())
			if err != nil {
				c.log(logPairs{"cmd": "LIST", "err": err})
				c.writeln("426 connection closed: transfer aborted")
			}
		}
	} else {
		_, err = fmt.Fprint(w, filename, c.lineEnding())
		if err != nil {
			c.log(logPairs{"cmd": "LIST", "err": err})
			c.writeln("426 connection closed: transfer aborted")
		}
	}
	c.writeln("226 closing data connection. list successful")
}

func (c *Conn) writeln(s ...interface{}) {
	if c.cmdErr != nil {
		return
	}
	s = append(s, "\r\n")
	_, c.cmdErr = fmt.Fprint(c.rw, s...)
}

func (c *Conn) lineEnding() string {
	if c.binary {
		return "\n"
	}
	return "\r\n"
}

// CmdErr returns command errors
func (c *Conn) CmdErr() error {
	return c.cmdErr
}

// Close closes the connection
func (c *Conn) Close() error {
	err := c.rw.Close()
	if err != nil {
		c.log(logPairs{"err": fmt.Errorf("closing command connection: %s", err)})
	}
	return err
}

func (c *Conn) pasv(args []string) {
	if len(args) > 0 {
		c.writeln("501 too many arguments")
		return
	}
	var firstError error
	storeFirstError := func(err error) {
		if firstError == nil {
			firstError = err
		}
	}

	var err error
	c.pasvListener, err = net.Listen("tcp4", "")
	storeFirstError(err)
	_, port, err := net.SplitHostPort(c.pasvListener.Addr().String())
	storeFirstError(err)
	ip, _, err := net.SplitHostPort(c.rw.LocalAddr().String())
	storeFirstError(err)
	addr, err := hostPortToFTP(fmt.Sprintf("%s:%s", ip, port))
	storeFirstError(err)
	if firstError != nil {
		c.pasvListener.Close()
		c.pasvListener = nil
		c.log(logPairs{"cmd": "PASV", "err": err})
		c.writeln("451 requested action aborted. local error in processing")
		return
	}
	c.writeln(fmt.Sprintf("227 =%s", addr))
}

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.writeln("501 usage: PORT a,b,c,d,p1,p2")
		return
	}
	var err error
	c.dataHostPort, err = hostPortFromFTP(args[0])
	if err != nil {
		c.log(logPairs{"cmd": "PORT", "err": err})
		c.writeln("501 can't parse address")
		return
	}
	c.writeln("200 PORT command successful")
}

// typ is short for type but type is a keyword
func (c *Conn) typ(args []string) {
	if len(args) < 1 || len(args) > 2 {
		c.writeln("501 usage: type takes 1 or 2 arguments")
		return
	}
	switch strings.ToUpper(strings.Join(args, " ")) {
	case "A", "A N":
		c.binary = false
	case "I", "L 8":
		c.binary = true
	default:
		c.writeln("504 unsupported type. supported types: A, A N, I, L 8")
		return
	}
	c.writeln("200 TYPE set")
}

func (c *Conn) stru(args []string) {
	if len(args) != 1 {
		c.writeln("501 usage: STRU F")
		return
	}
	if args[0] != "F" {
		c.writeln("504 only file structure is supported")
		return
	}
	c.writeln("200 STRU set")
}

func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.writeln("501 usage: RETR filename")
		return
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		c.log(logPairs{"cmd": "RETR", "err": err})
		c.writeln("550 file not found")
		return
	}
	c.writeln("150 file ok. sending")
	conn, err := c.dataConn()
	if err != nil {
		c.writeln("425 can't open data connection")
		return
	}
	defer conn.Close()
	if c.binary {
		_, err := io.Copy(conn, file)
		if err != nil {
			c.log(logPairs{"cmd": "RETR", "err": err})
			c.writeln("450 file unavailable")
			return
		}
	} else {
		// convert line endings from LF -> CRLF
		r := bufio.NewReader(file)
		w := bufio.NewWriter(conn)
		for {
			line, isPrefix, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				c.log(logPairs{"cmd": "RETR", "err": err})
				c.writeln("450 file unavailable")
				return
			}
			w.Write(line)
			if !isPrefix {
				w.Write([]byte("\r\n"))
			}
		}
		w.Flush()
	}
	c.writeln("226 transfer complete")
}
