package internal

import (
	"github.com/pmiguel/kiwi/pkg/protocol"
	"log"
	"net"
)

type Session struct {
	conn       net.Conn
	dispatcher *Dispatcher
}

const inboundBufferSize = 64

func (s *Session) StartSessionListener() {
	log.Printf("Started session listener for client: %s", s.conn.RemoteAddr().String())

	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		length, readErr := s.conn.Read(inboundBuffer)

		log.Printf("Read %d bytes", length)

		var response protocol.Response

		if readErr != nil {
			log.Printf("Error reading socket: %s. Closing session...", readErr)
			break
		}

		request, decodeErr := protocol.Decode[protocol.Request](inboundBuffer)

		if decodeErr != nil {
			log.Printf("Error decoding request %s", decodeErr)

			response = protocol.NewResponse("KIWI_DECODE_ERROR", true)
		} else {
			log.Printf("Processing %s %s %s", request.Command, request.Key, request.Value)
			response = s.dispatcher.Dispatch(request)
		}

		responseBytes, encodeErr := protocol.Encode[protocol.Response](&response)

		if encodeErr != nil {
			log.Printf("Error encoding response: %s", encodeErr)
			break
		}

		_, writeErr := s.conn.Write(responseBytes)

		if writeErr != nil {
			log.Printf("Error writing to socket: %s", writeErr)
			break
		}
	}

	if closeErr := s.conn.Close(); closeErr != nil {
		log.Printf("Error closing socket: %s", closeErr)
	}
}
