package ctci_chapter2

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in []int
		want bool
	}{
		{[]int {1,2,1,1,3,4,4,3,1,1,2,1}, true},
		{[]int {1,2,1,1,3,4,5,4,3,1,1,2,1}, true},
		{[]int {1,2,1,1,3,4,4,3,2,1}, false},
		{[]int {1,2,1,1,3,4,4,3,1,1,2}, false},
		{[]int {1,2,1,1,3,4,3,3,1,1,2,1}, false},
	}

	for _, c := range cases {
		inNode := NodeSetupSmart(c.in)
		got := inNode.IsPalindrome()
		if got != c.want {
			t.Errorf("IsPalindrome(%q) == %t, want %t", inNode.ListString(), got, c.want)
		}
	}
}