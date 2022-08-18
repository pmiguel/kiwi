package main

type CommandResult struct {
	Err    *string
	Result any
}

type HandleCommand interface {
	Execute(...any) *CommandResult
}

type Mediator struct {
	commandMap map[string]HandleCommand
}

func NewMediator() *Mediator {
	var m Mediator
	m.commandMap = make(map[string]HandleCommand)
	return &m
}

func (m *Mediator) RegisterCommand(command string, handler HandleCommand) {
	m.commandMap[command] = handler
}

func (m *Mediator) Execute(command string, args ...any) (*CommandResult, any) {
	handler, ok := m.commandMap[command]
	if !ok {
		return nil, "Unable to find command"
	}

	return handler.Execute(args), nil
}

type PingCommand struct {
}

func (p PingCommand) Execute(a ...any) *CommandResult {
	return &CommandResult{Err: nil, Result: "PONG"}
}
