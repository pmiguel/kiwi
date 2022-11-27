package internal

import (
	"github.com/pmiguel/kiwi/pkg/protocol"
	"strings"
)

type Dispatcher struct {
	storageManager *StorageManager
	server         *Server
	commands       map[string]Command
}

func NewDispatcher(server *Server, storageManager *StorageManager) *Dispatcher {
	dispatcher := &Dispatcher{server: server, storageManager: storageManager, commands: make(map[string]Command)}

	dispatcher.loadCommands()
	server.Dispatcher = dispatcher
	return dispatcher
}

func (d Dispatcher) loadCommands() {
	d.commands["PING"] = PingCommand{}
	d.commands["GET"] = GetCommand{storageManager: d.storageManager}
}

func (d Dispatcher) Dispatch(request *protocol.Request) protocol.Response {
	cmdUpper := strings.ToUpper(request.Command)

	cmd, ok := d.commands[cmdUpper]

	if ok {
		res, err := cmd.Exec(request.Arguments)

		if err != nil {
			return protocol.Response{Err: true, Content: err.Error()}
		}
		return protocol.Response{Err: false, Content: string(res)}
	}

	if cmdUpper == "SET" {
		d.storageManager.Set(request.Arguments[0], request.Arguments[1])
		return protocol.Response{Err: false, Content: "OK"}
	}

	if cmdUpper == "DEL" {
		d.storageManager.Delete(request.Arguments[0])

		return protocol.Response{Err: false, Content: "OK"}
	}

	if cmdUpper == "KEYS" {
		value := d.storageManager.Keys()

		return protocol.Response{Err: false, Content: value}
	}

	return protocol.Response{Err: true, Content: "KIWI_UNKNOWN_COMMAND"}
}
