package markhttp

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Handler interface {
	ServerHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
	Write([]byte) (int, error)
}

type responseWriter struct {
	Conn net.Conn
}

func (r *responseWriter) Write(bytes []byte) (int, error) {
	defer r.Conn.Close()
	// TODO: Status-Line = HTTP-Version SP Status-Code SP Reason-Phrase CRLF
	// TODO: Create struct to encapsulate response
	statusLine := "HTTP/1.1 200 ok\r\n"
	contentTypeHeader := "Content-Type: text/plain; charset=utf-8\r\n"
	contentLengthHeader := fmt.Sprintf("Content-Length: %v\r\n", len(bytes))

	responseBody := statusLine + contentLengthHeader + contentTypeHeader + "\r\n"

	return io.WriteString(r.Conn, responseBody+string(bytes))
}

type ServeMux struct {
	// TODO: populate with route and its handler
}

var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	// TODO: Based on Request data, delegate to respective handler to handle
	// TODO: Should check if any router is registered in ServeMux, if none, return 404
	// == For Debug Purpose ==
	_, err := w.Write([]byte("Hello I am mark\r\n"))
	if err != nil {
		panic(err)
	}
}

func ListenAndServe(address string, handler Handler) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer ln.Close()

	// Support HTTP/1.0 and only 1 connection at a time
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		// Prepare ResponseWriter & Request
		response := responseWriter{
			Conn: conn,
		}
		request := Request{}

		// TODO: Parse HTTP Request and populate Request struct
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
		}

		// Handle routes
		if handler != nil {
			handler.ServerHTTP(&response, &request)
		} else {
			// TODO: DefaultServeMux
			DefaultServeMux.ServeHTTP(&response, &request)
		}
	}

}
