package main

import (
	"github.com/pmiguel/kiwi/internal"
)

const (
	HOST = "0.0.0.0"
	PORT = "7170"
)

const (
	Server         int = 0
	StorageManager     = 1
	SessionManager     = 2
	Dispatcher         = 3
)

func main() {
	m := make(map[int]any)

	serverInstance := internal.NewServer(HOST, PORT)
	storageManager := internal.NewStorageManager(serverInstance)
	sessionManager := internal.NewSessionManager(serverInstance)
	dispatcher := internal.NewDispatcher(serverInstance, storageManager)

	m[Server] = serverInstance
	m[StorageManager] = storageManager
	m[SessionManager] = sessionManager
	m[Dispatcher] = dispatcher

	serverInstance.Start()
}
