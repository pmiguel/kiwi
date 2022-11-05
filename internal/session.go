package internal

import (
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"net"
)

type Session struct {
	conn       net.Conn
	dispatcher *Dispatcher
}

const inboundBufferSize = 64

func (s *Session) StartSessionListener() {
	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		_, readErr := s.conn.Read(inboundBuffer)

		var response protocol.Response

		if readErr != nil {
			fmt.Printf("Error reading socket: %s", readErr)
			break
		}

		request, decodeErr := protocol.Decode[protocol.Request](inboundBuffer)

		if decodeErr != nil {
			response = protocol.NewResponse("KIWI_DECODE_ERROR", true)
		} else {
			response = s.dispatcher.Dispatch(request)
		}

		responseBytes, encodeErr := protocol.Encode[protocol.Response](&response)

		if encodeErr != nil {
			fmt.Printf("Error encoding response: %s", encodeErr)
			break
		}

		_, writeErr := s.conn.Write(responseBytes)

		if writeErr != nil {
			fmt.Printf("Error writing to socket: %s", writeErr)
			break
		}
	}

	err := s.conn.Close()
	if err != nil {
		fmt.Printf("Error closing socket: %s", err)
	}
}
