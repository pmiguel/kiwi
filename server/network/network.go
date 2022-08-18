package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pmiguel/kiwi/common/protocol"
	"log"
	"net"
)

const proto = "tcp"
const inboundBufferSize = 64

type Server struct {
	host    string
	port    string
	running bool
}

func NewServer(host string, port string) *Server {
	return &Server{
		host:    host,
		port:    port,
		running: false,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen(proto, s.host+":"+s.port)
	s.running = true

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Kiwi listening on port " + s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleIncomingRequest(conn)
	}
	listener.Close()
}

func handleIncomingRequest(conn net.Conn) {
	sender := conn.RemoteAddr().String()
	fmt.Println("<= " + sender)
	counter := 0
	for {
		packet := make([]byte, inboundBufferSize)
		_, err := conn.Read(packet)

		if err != nil {
			break
		}

		var buffer bytes.Buffer
		buffer.Write(packet)

		dec := gob.NewDecoder(&buffer)

		var req protocol.Request
		err = dec.Decode(&req)

		if err == nil {
			fmt.Printf("<< 0x%x (%d bytes) n:%d {%s}\n", buffer, buffer.Len(), counter, sender)
			fmt.Printf("\t<< %s", req.String())
		} else {
			fmt.Printf("<< %s", err)
		}

		counter++
		res := protocol.Response{Content: "PONG"}

		conn.Write(res.Bytes())
	}
	fmt.Println("=> " + sender)
	conn.Close()
}
