package internal

import (
	"github.com/pmiguel/kiwi/pkg/protocol"
	"github.com/pmiguel/kiwi/pkg/protocol/resp"
)

func ParseRawRequest(input []byte) protocol.Request {
	respArray := resp.DecodeArray(string(input))

	if len(respArray) > 1 {
		return protocol.Request{Command: respArray[0], Arguments: respArray[1:]}
	}

	return protocol.Request{Command: respArray[0]}
}
