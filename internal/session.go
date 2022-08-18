package internal

import (
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"net"
)

type Session struct {
	conn net.Conn
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
	return protocol.Response{Err: true, Content: "KIWI_UNSUPPORTED_COMMAND"}
}
