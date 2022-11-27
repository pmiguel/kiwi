package internal

type CommandResult string
type CommandArguments []string

type Command interface {
	Exec(arguments CommandArguments) (CommandResult, error)
}
