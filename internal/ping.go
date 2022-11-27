package internal

type PingCommand struct{}

func (c PingCommand) Exec(_ CommandArguments) (CommandResult, error) {
	return "PONG", nil
}
