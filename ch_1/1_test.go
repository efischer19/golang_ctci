 package ctci_chapter1

import "testing"

func testAllUnique(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abbc", false},
		{"ABC123ABC123", false},
	}

	for _, c := range cases {
		got1 := AllUniqueNaive(c.in)
		if got1 != c.want {
			t.Errorf("AllUniqueNaive(%q) == %q, want %q", c.in, got1, c.want)
		}
		got2 = AllUniqueSet(c.in)
		if got2 != c.want {
			t.Errorf("AllUniqueNaive(%q) == %q", want %q, c.in, got2, c.want)
		}
	}
}
