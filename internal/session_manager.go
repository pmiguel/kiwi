package internal

import (
	"log"
	"net"
)

type SessionManager struct {
	sessions map[net.Addr]*Session
	server   *Server
}

func NewSessionManager(server *Server) *SessionManager {
	sessionManager := &SessionManager{
		sessions: make(map[net.Addr]*Session),
		server:   server,
	}

	server.SessionManager = sessionManager
	return sessionManager
}

func (sm *SessionManager) RegisterSession(conn net.Conn) {
	session := NewSession(conn, sm.server.Dispatcher)

	sm.sessions[conn.RemoteAddr()] = &session

	log.Printf("Registered session for client: %s", conn.RemoteAddr().String())

	go session.StartSessionListener()
}

func (sm *SessionManager) GetSession(addr net.Addr) *Session {
	return sm.sessions[addr]
}
