package main

import (
	"github.com/pmiguel/kiwi/internal"
)

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	server := internal.NewServer(HOST, PORT)
	server.Start()
}
