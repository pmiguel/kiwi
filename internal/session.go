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
	conn := s.conn
	sender := conn.RemoteAddr().String()

	for {
		inboundBuffer := make([]byte, inboundBufferSize)
		_, err := conn.Read(inboundBuffer)

		if err != nil {
			break
		}

		request, err := protocol.Decode[protocol.Request](inboundBuffer)

		var response protocol.Response

		if err != nil {
			response = protocol.NewResponse("KIWI_DECODE_ERROR", true)
		} else {
			response = s.dispatcher.Dispatch(request)
		}

		responseBytes, err := protocol.Encode[protocol.Response](&response)

		conn.Write(responseBytes)
	}

	fmt.Println("=> " + sender)
	conn.Close()
}
