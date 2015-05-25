package ctci_chapter1

import "testing"

func TestCompress(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"aabcccccaaa","a2b1c5a3"},
		{"abbc", "abbc"},
		{"aaAAAAzuDDDwww", "a2A4z1u1D3w3"},
	}

	for _, c := range cases {
		got := Compress(c.in)
		if got != c.want {
			t.Errorf("SpaceEncode(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
