package resp

import "fmt"

const prefixSimpleString = '+'

type SimpleString string

func (str SimpleString) Encode() string {
	return fmt.Sprintf("%c%s%s", prefixSimpleString, str, sepMarker)
}

func DecodeSimpleString(input string) SimpleString {
	chars := []rune(input)

	if chars[0] != prefixSimpleString {
		return ""
	}

	buf := make([]rune, len(chars)-3)
	for i := 1; chars[i] != '\r' && i < len(chars)-1; i++ {
		buf[i-1] = chars[i]
	}

	return SimpleString(buf)
}
