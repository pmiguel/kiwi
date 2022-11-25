package main

import (
	"bufio"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
	"github.com/pmiguel/kiwi/pkg/protocol/kcp"
	"github.com/pmiguel/kiwi/pkg/protocol/resp"
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

	prompt := conn.RemoteAddr().String() + " >> "

	for {

		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')

		rawInputString := strings.TrimSpace(text)

		tokens := strings.Split(rawInputString, " ")

		var respArray = resp.Array{}

		for i := 0; i < len(tokens); i++ {
			respArray = append(respArray, resp.BulkString(tokens[i]))
		}

		// writeBytes, _ := kcp.Encode[protocol.Request](protocol.NewRequest(, "#", "#"))

		_, writeErr := conn.Write([]byte(respArray.Encode()))

		if writeErr != nil {
			log.Fatal(writeErr)
		}

		readBuffer := make([]byte, 1024)
		_, readErr := conn.Read(readBuffer)
		if readErr != nil {
			log.Fatal(readErr)
		}

		dec, _ := kcp.Decode[protocol.Response](readBuffer)

		fmt.Printf("%s (Err: %t)\n", dec.Content, dec.Err)
	}
	conn.Close()
}
