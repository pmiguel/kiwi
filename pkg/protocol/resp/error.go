package resp

import "fmt"

const prefixError = '-'

type Error string

func (str Error) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixError, str, sepMarker)
}

func DecodeError(input string) Error {
	chars := []rune(input)

	if chars[0] != prefixError {
		return ""
	}

	buf := make([]rune, len(chars)-3)
	for i := 1; chars[i] != '\r' && i < len(chars)-1; i++ {
		buf[i-1] = chars[i]
	}
	return Error(buf)
}
