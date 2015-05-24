 package ctci_chapter1

import "testing"

func TestIsPerm(t *testing.T) {
	cases := []struct {
		in1 string
		in2 string
		want bool
	}{
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "ZYXWVUTSRQPONMLKJIHGFEDCBAzyxvwutsrqponmlkjihgfedcba", true},
		{"abbc", "cabb", true},
		{"ABC123ABC123", "123123ABCABC", true},
		{"abbcdef", "abccdef", false},
	}

	for _, c := range cases {
		got := IsPermutation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("IsPermutation(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}
