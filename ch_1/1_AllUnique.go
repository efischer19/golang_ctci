package ctci_chapter1

func AllUniqueNaive(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func AllUniqueSet(s string) bool {
	var set = make(map[rune]bool)
	for _, c := range s {
		if set[c] {
			return false
		}
		set[c] = true
	}
	return true
}
