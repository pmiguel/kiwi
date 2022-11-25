package resp

const sepMarker = "\r\n"

type Resp interface {
	Encode() string
}

func DecodeLength(input []rune, startIndex int) (length int, currentIndex int) {
	length = 0
	signal := 1
	currentIndex = startIndex
	for input[currentIndex] != '\r' {
		if currentIndex == 0 && input[currentIndex] == '-' {
			signal = -(signal)
		} else {
			length = (length * 10) + (int(input[currentIndex]) - int('0'))
		}
		currentIndex++
	}
	return length * signal, currentIndex + 2 // Skip \r\n
}
