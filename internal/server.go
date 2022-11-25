package internal

import (
	"log"
	"net"
)

const (
	Host     = "0.0.0.0"
	Port     = "7170"
	Protocol = "tcp"
)

type Server struct {
	host           string
	port           string
	running        bool
	SessionManager *SessionManager
	StorageManager *StorageManager
	Dispatcher     *Dispatcher
}

func (s *Server) Start() {

	if s.SessionManager == nil {
		log.Fatal("Unable to initiate server. Session manager not defined.")
	}

	if s.StorageManager == nil {
		log.Fatal("Unable to initiate server. Storage manager not defined.")
	}

	listener, err := net.Listen(Protocol, s.host+":"+s.port)
	s.running = true

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on port " + s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Client connected: %s", conn.RemoteAddr().String())

		s.SessionManager.RegisterSession(conn)
	}
	listener.Close()
}

func NewServer(host string, port string) *Server {
	return &Server{
		host:    host,
		port:    port,
		running: false,
	}
}
