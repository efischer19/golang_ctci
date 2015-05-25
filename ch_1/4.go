package ctci_chapter1

import "bytes"

func SpaceEncode(input string) string {
	var buffer bytes.Buffer
	var ready,writeSpace = false,false
	for _, c := range input {
		if c != ' ' {
			if writeSpace {
				buffer.WriteString("%20")
			}
			buffer.WriteString(string(c))
			ready = true
			writeSpace = false
		} else if ready {
			writeSpace = true
		}
	}
	return buffer.String()
}
