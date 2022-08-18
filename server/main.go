package main

import "fmt"

const (
	HOST = "localhost"
	PORT = "7170"
)

func main() {
	// server := NewServer(HOST, PORT)

	mediator := NewMediator()

	mediator.RegisterCommand("PING", PingCommand{})

	res, err := mediator.Execute("PING")
	_, err2 := mediator.Execute("UNAVAILABLE")

	fmt.Println(res.Result)
	fmt.Println(err)

	fmt.Println(err2)

	// server.Start()
}
