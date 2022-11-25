package main

import (
	"bufio"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"github.com/pmiguel/kiwi/pkg/protocol/kcp"
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
		var readBuffer = make([]byte, 1024)

		fmt.Print(conn.RemoteAddr().String() + " >> ")
		text, _ := reader.ReadString('\n')

		target := strings.TrimSpace(text)

		tokens := strings.Split(target, " ")

		var request *protocol.Request

		if len(tokens) == 1 {
			request = protocol.NewRequest(tokens[0], "#", "#")
		}

		if len(tokens) == 2 {
			request = protocol.NewRequest(tokens[0], tokens[1], "#")
		}

		if len(tokens) == 3 {
			request = protocol.NewRequest(tokens[0], tokens[1], tokens[2])
		}

		requestBytes, _ := kcp.Encode[protocol.Request](request)

		conn.Write(requestBytes)
		conn.Read(readBuffer)
		dec, _ := kcp.Decode[protocol.Response](readBuffer)

		fmt.Printf("%s (Err: %t)\n", dec.Content, dec.Err)
	}
	conn.Close()
}
