package internal

import (
	"github.com/google/uuid"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"log"
	"net"
)

type Session struct {
	id         uuid.UUID
	conn       net.Conn
	dispatcher *Dispatcher
}

func NewSession(conn net.Conn, dispatcher *Dispatcher) Session {
	return Session{
		id:         uuid.New(),
		conn:       conn,
		dispatcher: dispatcher,
	}
}

const inboundBufferSize = 64

func (s *Session) StartSessionListener() {
	sessionId := s.id.String()
	log.Printf("%s Started session listener for client: %s", sessionId, s.conn.RemoteAddr().String())

	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		length, readErr := s.conn.Read(inboundBuffer)

		log.Printf("%s Read %d bytes", sessionId, length)

		var response protocol.Response

		if readErr != nil {
			log.Printf("%s Error reading socket: %s. Closing session...", sessionId, readErr)
			break
		}

		request, decodeErr := protocol.Decode[protocol.Request](inboundBuffer)

		if decodeErr != nil {
			log.Printf("%s Error decoding request %s", sessionId, decodeErr)

			response = protocol.NewResponse("KIWI_DECODE_ERROR", true)
		} else {
			log.Printf("%s Processing %s %s %s", sessionId, request.Command, request.Key, request.Value)
			response = s.dispatcher.Dispatch(request)
		}

		responseBytes, encodeErr := protocol.Encode[protocol.Response](&response)

		if encodeErr != nil {
			log.Printf("%s Error encoding response: %s", sessionId, encodeErr)
			break
		}

		_, writeErr := s.conn.Write(responseBytes)

		if writeErr != nil {
			log.Printf("%s Error writing to socket: %s", sessionId, writeErr)
			break
		}
	}

	if closeErr := s.conn.Close(); closeErr != nil {
		log.Printf("%s Error closing socket: %s", sessionId, closeErr)
	}
}
