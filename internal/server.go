package internal

import (
	"fmt"
	"log"
	"net"
)

const proto = "tcp"
const inboundBufferSize = 64

type Server struct {
	host           string
	port           string
	running        bool
	sessionManager *SessionManager
}

func NewServer(host string, port string) *Server {
	return &Server{
		host:           host,
		port:           port,
		running:        false,
		sessionManager: NewSessionManager(),
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
		session := s.sessionManager.RegisterSession(conn)
		go session.StartSessionListener()
	}
	listener.Close()
}
