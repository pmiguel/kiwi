package protocol

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Response struct {
	Err     bool
	Content string
}

func (r Response) Bytes() []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(r)

	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
