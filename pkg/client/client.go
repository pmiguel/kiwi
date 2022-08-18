package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
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

func decodeResponse(packet []byte) (protocol.Response, error) {
	buffer := bytes.NewBuffer(packet)
	dec := gob.NewDecoder(buffer)

	var res protocol.Response
	err := dec.Decode(&res)

	return res, err
}

func encodeRequest(request *protocol.Request) ([]byte, error) {
	buffer := bytes.Buffer{}
	dec := gob.NewEncoder(&buffer)

	err := dec.Encode(request)
	return buffer.Bytes(), err
}

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

		requestBytes, _ := encodeRequest(protocol.NewRequest(target, "key", "value"))
		conn.Write(requestBytes)
		conn.Read(readBuffer)
		dec, _ := decodeResponse(readBuffer)

		fmt.Printf("%s (Err: %t)\n", dec.Content, dec.Err)
	}
	conn.Close()
}
