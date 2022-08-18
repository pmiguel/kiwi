package protocol

import (
	"bytes"
	"fmt"
)

type Request struct {
	Command string
	Key     string
	Value   string
}

func NewRequest(command string, key string, value string) *Request {
	return &Request{Command: command, Key: key, Value: value}
}

func (r *Request) String() string {
	return fmt.Sprintf("[%s] %s -> %s\n", r.Command, r.Key, r.Value)
}

func (r *Request) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanf(b, "%s %s %s\n", &r.Command, &r.Key, &r.Value)
	return err
}

func (r *Request) MarshalBinary() (data []byte, err error) {
	var b bytes.Buffer
	_, err = fmt.Fprintf(&b, "%s %s %s\n", r.Command, r.Key, r.Value)
	return b.Bytes(), err
}
