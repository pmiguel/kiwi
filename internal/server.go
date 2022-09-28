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
	SessionManager *SessionManager
	StorageManager *StorageManager
	Dispatcher     *Dispatcher
}

func NewServer(host string, port string) *Server {
	return &Server{
		host:    host,
		port:    port,
		running: false,
	}
}

func (s *Server) Start() {

	if s.SessionManager == nil {
		panic("Unable to initiate server. Session manager not defined.")
	}

	if s.StorageManager == nil {
		panic("Unable to initiate server. Storage manager not defined.")
	}

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
		s.SessionManager.RegisterSession(conn)
	}
	listener.Close()
}
