package internal

import (
	"fmt"
	"net"
)

type SessionManager struct {
	sessions          map[net.Addr]*Session
	autoStartSessions bool
}

func NewSessionManager(autoStartSession bool) *SessionManager {
	return &SessionManager{
		sessions:          make(map[net.Addr]*Session),
		autoStartSessions: autoStartSession,
	}
}

func (sm *SessionManager) RegisterSession(conn net.Conn) {
	session := NewSession(conn)
	sm.sessions[conn.RemoteAddr()] = session
	fmt.Println("<= " + conn.RemoteAddr().String())

	if sm.autoStartSessions {
		go session.StartSessionListener()
	}
}

func (sm *SessionManager) GetSession(addr net.Addr) *Session {
	return sm.sessions[addr]
}
