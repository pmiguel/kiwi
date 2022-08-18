package internal

import (
	"fmt"
	"log"
	"net"
)

const proto = "tcp"

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
		sessionManager: NewSessionManager(true),
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
		s.sessionManager.RegisterSession(conn)
	}
	listener.Close()
}
