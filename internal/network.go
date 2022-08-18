package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
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
	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		length, err := conn.Read(inboundBuffer)

		if err != nil {
			break
		}

		request, decodeError := decodeRequest(inboundBuffer)

		if decodeError == nil {
			fmt.Printf("<< 0x%x (%d bytes) {%s}\n", inboundBuffer, length, sender)
			fmt.Printf("\t<< %s", request.String())
		} else {
			fmt.Printf("<< %s", err)
		}

		response := executeCommand(&request)
		responseBytes, err := encodeResponse(&response)

		conn.Write(responseBytes)
	}
	fmt.Println("=> " + sender)
	conn.Close()
}

func decodeRequest(packet []byte) (protocol.Request, error) {
	buffer := bytes.NewBuffer(packet)
	dec := gob.NewDecoder(buffer)

	var req protocol.Request
	err := dec.Decode(&req)

	return req, err
}

func encodeResponse(response *protocol.Response) ([]byte, error) {
	buffer := bytes.Buffer{}
	dec := gob.NewEncoder(&buffer)

	err := dec.Encode(response)
	return buffer.Bytes(), err
}

func executeCommand(request *protocol.Request) protocol.Response {
	if request.Command == "PING" {
		return protocol.Response{Err: false, Content: "PONG"}
	}
	return protocol.Response{Err: true, Content: "KIWI_UNSUPPORTED_COMMAND"}
}
