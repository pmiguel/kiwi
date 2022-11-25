package resp

import "fmt"

const prefixInteger = ':'

type Int int64

func (i Int) Encode() string {
	return fmt.Sprintf("%c%d%s", prefixInteger, i, sepMarker)
}

func DecodeInt(input string) Int {
	var out Int
	_, err := fmt.Sscanf(input, ":%d\r\n", &out)
	if err != nil {
		return Int(0)
	}
	return Int(out)
}
