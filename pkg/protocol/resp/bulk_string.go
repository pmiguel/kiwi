package resp

import "fmt"

const prefixBulkString = '$'

type BulkString string

func (str BulkString) Encode() string {
	return fmt.Sprintf("%c%d%s%s%s", prefixBulkString, len(str), sepMarker, str, sepMarker)
}

func DecodeBulkString(input string) BulkString {
	chars := []rune(input)

	if chars[0] != prefixBulkString {
		return ""
	}

	length, index := DecodeLength(chars, 1)
	buff := make([]rune, length)
	for i := 0; i < length; i++ {
		buff[i] = chars[index+i]
	}
	return BulkString(buff)
}
