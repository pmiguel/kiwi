package internal

import (
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"net"
)

type Session struct {
	conn net.Conn
}

const inboundBufferSize = 64

func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

func (s *Session) StartSessionListener() {
	conn := s.conn
	sender := conn.RemoteAddr().String()

	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		length, err := conn.Read(inboundBuffer)

		if err != nil {
			break
		}

		request, err := protocol.Decode[protocol.Request](inboundBuffer)

		if err == nil {
			fmt.Printf("<< 0x%x (%d bytes) {%s}\n", inboundBuffer, length, sender)
			fmt.Printf("\t<< %s", request.String())
		} else {
			fmt.Printf("<< %s", err)
		}

		response := s.executeCommand(&request)
		responseBytes, err := protocol.Encode[protocol.Response](&response)

		conn.Write(responseBytes)
	}

	fmt.Println("=> " + sender)
	conn.Close()
}

func (s *Session) executeCommand(request *protocol.Request) protocol.Response {
	if request.Command == "PING" {
		return protocol.Response{Err: false, Content: "PONG"}
	}

	if request.Command == "SET" {
		fmt.Printf("\t%s %s %s\n", request.Command, request.Key, request.Value)
	}

	return protocol.Response{Err: true, Content: "KIWI_UNSUPPORTED_COMMAND"}
}
