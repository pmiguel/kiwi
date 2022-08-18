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
	server.Start()
}
