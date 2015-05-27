 package ctci_chapter1

import "testing"

func TestSpaceEncode(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"Hello, World!", "Hello,%20World!"},
		{"abbc", "abbc"},
		{"  me ow  ", "me%20ow"},
		{"Mr John Smith   ", "Mr%20John%20Smith"},
	}

	for _, c := range cases {
		got := SpaceEncode(c.in)
		if got != c.want {
			t.Errorf("SpaceEncode(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
