package main

import (
	server "github.com/pmiguel/kiwi/server/network"
)

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	server := server.New(HOST, PORT)
	server.Start()
}
