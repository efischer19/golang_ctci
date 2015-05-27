 package ctci_chapter1

import "testing"

func TestIsRotation(t *testing.T) {
	cases := []struct {
		in1 string
		in2 string
		want bool
	}{
		{"Hello, World!", "World!Hello, ", true},
		{"erbottlewat","waterbottle", true},
		{"meowmeowmeow", "meowmeow", false},
		{"abcd", "efgh", false},
		{"","",true},
	}

	for _, c := range cases {
		got := IsRotation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("IsRotation(%q,%q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}
