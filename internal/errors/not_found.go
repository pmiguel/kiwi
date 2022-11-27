package errors

type KeyNotFoundError string

func (e KeyNotFoundError) Error() string {
	return "KIWI_KEY_NOT_FOUND"
}
