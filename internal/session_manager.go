package internal

import (
	"log"
	"net"
)

type Manager struct {
	sessions          map[net.Addr]*Session
	autoStartSessions bool
	server            *Server
}

func NewSessionManager(server *Server) *Manager {
	sessionManager := &Manager{
		sessions: make(map[net.Addr]*Session),
		server:   server,
	}

	server.SessionManager = sessionManager
	return sessionManager
}

func (sm *Manager) RegisterSession(conn net.Conn) {
	session := NewSession(conn, sm.server.Dispatcher)

	sm.sessions[conn.RemoteAddr()] = &session

	log.Printf("Registered session for client: %s", conn.RemoteAddr().String())

	go session.StartSessionListener()
}

func (sm *Manager) GetSession(addr net.Addr) *Session {
	return sm.sessions[addr]
}
