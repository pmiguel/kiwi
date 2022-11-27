package internal

import (
	"fmt"
	"github.com/pmiguel/kiwi/internal/errors"
)

type GetCommand struct {
	storageManager *StorageManager
}

func (c GetCommand) Exec(arguments CommandArguments) (CommandResult, error) {
	value := c.storageManager.Get(arguments[0])

	if value == nil {
		return "", errors.KeyNotFoundError(arguments[0])
	}

	return CommandResult(fmt.Sprintf("%v", value)), nil
}
