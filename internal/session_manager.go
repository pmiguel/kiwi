package internal

import (
	"fmt"
	"net"
)

type SessionManager struct {
	sessions map[net.Addr]Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[net.Addr]Session),
	}
}

func (sm *SessionManager) RegisterSession(conn net.Conn) *Session {
	session := Session{conn: conn}
	sm.sessions[conn.RemoteAddr()] = session
	fmt.Println("<= " + conn.RemoteAddr().String())
	return &session
}
