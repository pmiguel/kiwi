package internal

import (
	"fmt"
	"github.com/pmiguel/kiwi/pkg/protocol"
)

type Dispatcher struct {
	storageManager *StorageManager
	server         *Server
}

func NewDispatcher(server *Server, storageManager *StorageManager) *Dispatcher {
	dispatcher := &Dispatcher{server: server, storageManager: storageManager}

	server.Dispatcher = dispatcher
	return dispatcher
}

func (d *Dispatcher) Dispatch(request *protocol.Request) protocol.Response {
	if request.Command == "PING" {
		return protocol.Response{Err: false, Content: "PONG"}
	}

	if request.Command == "SET" {
		d.storageManager.Set(request.Key, request.Value)
		return protocol.Response{Err: false, Content: "OK"}
	}

	if request.Command == "GET" {
		value := d.storageManager.Get(request.Key)

		if value == nil {
			return protocol.Response{Err: true, Content: "KIWI_KEY_NOT_FOUND"}
		}

		return protocol.Response{Err: false, Content: fmt.Sprintf("%v", value)}
	}

	if request.Command == "DEL" {
		d.storageManager.Delete(request.Key)

		return protocol.Response{Err: false, Content: "OK"}
	}

	if request.Command == "KEYS" {
		value := d.storageManager.Keys()

		return protocol.Response{Err: false, Content: value}
	}

	return protocol.Response{Err: true, Content: "KIWI_UNKNOWN_COMMAND"}
}
