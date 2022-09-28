package protocol

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Response struct {
	Err     bool
	Content string
}

func NewResponse(content string, error bool) Response {
	return Response{Content: content, Err: error}
}

func (r *Response) Bytes() []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(r)

	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func (r *Response) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanf(b, "%t %s\n", &r.Err, &r.Content)
	return err
}

func (r *Response) MarshalBinary() (data []byte, err error) {
	var b bytes.Buffer
	_, err = fmt.Fprintf(&b, "%t %s\n", r.Err, r.Content)
	return b.Bytes(), err
}
