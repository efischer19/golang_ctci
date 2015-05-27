package ctci_chapter1

import "strings"

//note: I'll be using strings.Contains() as my 'isSubstring' method as described in the book
func IsRotation(in1 string, in2 string) bool {
	if len(in1) != len(in2) {
		return false
	}
	var temp = in1+in1
	return strings.Contains(temp, in2)
}
