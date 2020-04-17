package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/0x00101010/simple_http_server/internal/connpool"
	"github.com/0x00101010/simple_http_server/internal/fs"
)

// Server represents a simple http server which supports only GET request
type Server struct {
	dir      string
	port     int
	listener net.Listener
	conns    connpool.ConnPool
}

const bufsize = 1024

var errResponse = `
HTTP/1.1 %[1]v
Content-Type: text/plain;charset=utf-8
Content-Length:
Date: 

%[1]v
`

var okResponse = `
HTTP/1.1 200 OK
Content-Type: %v;charset=utf-8
Content-Length: %v
Date: %v

`

// Init initializes server with correct state
func Init(dir string, port int) *Server {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Printf("Unable to start TCP listener: %v", err)
		os.Exit(1)
	}

	server := &Server{
		dir:      dir,
		port:     port,
		listener: listener,
		conns:    connpool.InitConnPool(),
	}
	return server
}

// ListenAndServe starts the http listener and start accepting incoming
// http requests
func (s *Server) ListenAndServe() {
	for {
		conn, err := s.conns.TryAcceptNewConn(s.listener)
		if err == connpool.ErrFullCapacity {
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}

		// one minute live time for persistent connections
		conn.SetDeadline(time.Now().Add(60 * time.Second))

		go s.HandleRequest(conn)
	}
}

// HandleRequest is the function for handling http request
func (s *Server) HandleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		req, err := http.ReadRequest(reader)

		// make sure request is well-formed
		if err != nil {
			fmt.Fprintf(conn, errResponse, "400 Bad Request")
			conn.Close()
			return
		}

		// make sure target file exists
		url := req.URL.Path
		if url == "/" {
			url = "/index.htm"
		}
		path := filepath.Join(s.dir, url)

		if !fs.Exists(path) {
			fmt.Fprintf(conn, errResponse, "404 Not Found")
			conn.Close()
			return
		}

		// make sure target file is ok for access
		if !fs.AllowRead(path) {
			fmt.Fprintf(conn, errResponse, "403 Forbidden")
			conn.Close()
			return
		}

		data, _ := ioutil.ReadFile(path)

		// write header response to client
		timeStr := time.Now().Format(time.RFC1123)
		fmt.Fprintf(conn, okResponse, contentType(path), len(data), timeStr)

		// write body
		conn.Write(data)
	}
}

func contentType(path string) string {
	ext := filepath.Ext(path)
	mime := mime.TypeByExtension(ext)
	if mime == "" {
		mime = "text/plain"
	}
	return mime
}
