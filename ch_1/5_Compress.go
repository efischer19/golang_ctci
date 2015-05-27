package ctci_chapter1

import (
	"bytes"
	"strconv"
)

func Compress(input string) string {
	var buffer bytes.Buffer
	var curRune rune = 0
	var curCount = 0
	for _, c := range input {
		if c != curRune {
			if curRune != 0 {
				buffer.WriteRune(curRune)
				buffer.WriteString(strconv.Itoa(curCount))
			}
			curRune = c
			curCount = 1
		} else {
			curCount++
		}
	}
	buffer.WriteRune(curRune)
	buffer.WriteString(strconv.Itoa(curCount))

	var ret = buffer.String()
	if len(ret) < len(input) {
		return ret
	}
	return input
}
