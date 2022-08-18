package main

import (
	"fmt"
	"github.com/pmiguel/kiwi/internal"
)

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	// server := NewServer(HOST, PORT)

	mediator := internal.NewMediator()

	mediator.RegisterCommand("PING", internal.PingCommand{})

	res, err := mediator.Execute("PING")
	_, err2 := mediator.Execute("UNAVAILABLE")

	fmt.Println(res.Result)
	fmt.Println(err)

	fmt.Println(err2)

	// server.Start()
}
