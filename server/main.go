package main

import "github.com/pmiguel/kiwi/server/network"

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	network.StartServer(HOST, PORT)
}
