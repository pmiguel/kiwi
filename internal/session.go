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

func (s *Session) read() (*protocol.Request, error) {
	var err error = nil
	var request protocol.Request

	inboundBuffer := make([]byte, inboundBufferSize)
	_, err = s.conn.Read(inboundBuffer)

	if err == nil {
		request, err = protocol.Decode[protocol.Request](inboundBuffer)
	}

	return &request, err
}

func (s *Session) write(response *protocol.Response) error {
	responseBytes, err := protocol.Encode[protocol.Response](response)

	if err == nil {
		_, err = s.conn.Write(responseBytes)
	}
	return err
}

func (s *Session) StartSessionListener() {
	sessionId := s.id.String()
	log.Printf("%s Started session listener for client: %s", sessionId, s.conn.RemoteAddr().String())

	for {
		var response protocol.Response

		if request, readErr := s.read(); readErr == nil {
			log.Printf("%s Processing request %s %s %s", sessionId, request.Command, request.Key, request.Value)
			response = s.dispatcher.Dispatch(request)
		} else {
			log.Printf("%s Error reading request. %s", sessionId, readErr)
			response = protocol.NewResponse("KIWI_DECODE_ERROR", true)
		}

		if writeErr := s.write(&response); writeErr != nil {
			log.Printf("%s Error writing response. Aborting session: %s", sessionId, writeErr)
			break
		}
	}

	if err := s.conn.Close(); err != nil {
		log.Printf("%s Error closing socket: %s", sessionId, err)
	}
}
