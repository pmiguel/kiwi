package internal

import (
	"github.com/google/uuid"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"github.com/pmiguel/kiwi/pkg/protocol/encoding"
	"log"
	"net"
)

const inboundBufferSize uint8 = 64

// Session represents an active connection to a connected client.
type Session struct {
	id         uuid.UUID
	conn       net.Conn
	dispatcher *Dispatcher
}

// NewSession creates a new Session struct containing the connection information,
// a reference to the internal.Dispatcher, along with a unique identifier for the connected
// client connection.
func NewSession(conn net.Conn, dispatcher *Dispatcher) Session {
	return Session{
		id:         uuid.New(),
		conn:       conn,
		dispatcher: dispatcher,
	}
}

// The read function reads inboundBufferSize bytes from the connected socket
// and returns a protocol.Request object representing the request being made to the server
// It returns either the protocol.Request in the expected format, or an error representing either a read error
// or a serialization error
func (s *Session) read() (*protocol.Request, error) {
	var err error = nil
	var request protocol.Request

	inboundBuffer := make([]byte, inboundBufferSize)
	_, err = s.conn.Read(inboundBuffer)

	if err == nil {
		request, err = encoding.Decode[protocol.Request](inboundBuffer)
	}

	return &request, err
}

// The write function receives a protocol.Response object containing the result to a given command
// and writes it into the connected buffer, as a byte array
func (s *Session) write(response *protocol.Response) error {
	responseBytes, err := encoding.Encode[protocol.Response](response)

	if err == nil {
		_, err = s.conn.Write(responseBytes)
	}
	return err
}

// StartSessionListener starts a new handler for the connected client that will receive protocol.Request and
// respond with protocol.Response containing the result of the operation.
// This operation happens continuously until the server or the client disconnects.
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
