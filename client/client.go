package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "7170"
	TYPE = "tcp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial(TYPE, HOST+":"+PORT)

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(strings.TrimSpace(text)))

		buffer := make([]byte, 1024)
		length, _ := conn.Read(buffer)
		fmt.Printf("<< %s \n", string(buffer[:length]))
	}
	conn.Close()
}
