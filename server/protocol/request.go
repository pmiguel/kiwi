package protocol

type Request struct {
	Command uint16
	Key     string
	Value   string
}
