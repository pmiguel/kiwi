package main

import (
	kiwi "github.com/pmiguel/kiwi/server/network"
)

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	server := kiwi.New(HOST, PORT)
	server.Start()
}
