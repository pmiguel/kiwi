package main

import (
	"github.com/pmiguel/kiwi/internal"
)

const (
	HOST = "0.0.0.0"
	PORT = "7170"
)

func main() {
	server := internal.NewServer(HOST, PORT)
	sessionManager := internal.NewSessionManager(server)
	storageManager := internal.NewStorageManager(server)
	dispatcher := internal.NewDispatcher(server, storageManager)

	server.StorageManager = storageManager
	server.SessionManager = sessionManager
	server.Dispatcher = dispatcher
	server.Start()
}
