package network

import (
	"fmt"
	"github.com/pmiguel/kiwi/server/protocol"
	"log"
	"net"
)

const proto = "tcp"
const inboundBufferSize = 1024

type Server struct {
	host    string
	port    string
	running bool
}

func New(host string, port string) *Server {
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
		buffer := make([]byte, inboundBufferSize)
		length, err := conn.Read(buffer)

		if err != nil {
			break
		}

		content := string(buffer[:length])
		fmt.Printf("<< %s (%d bytes) [0x%x], n:%d {%s}\n", content, length, content, counter, sender)
		counter++

		res := protocol.Response{Content: "PONG"}

		conn.Write(res.Bytes())
	}
	fmt.Println("=> " + sender)
	conn.Close()
}
