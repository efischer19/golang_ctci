package ctci_chapter1

func IsPermutation(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	set1 := make(map[rune]int)
	set2 := make(map[rune]int)
	for _, c := range first {
		set1[c]++
	}
	for _, c := range second {
		set2[c]++
	}

	if len(set1) != len(set2) {
		return false
	}

	for k,v := range set1 {
		if set2[k] != v {
			return false
		}
		delete(set2, k)
	}

	return len(set2) == 0
}
