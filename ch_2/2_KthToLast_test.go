package ctci_chapter2

import "testing"

func TestKthToLast(t *testing.T) {
	cases := []struct {
		inList []int
		inK int
		want int
	}{
		{[]int {1,2,2,3,4,2,3,6,7,2,4,1,8}, 3, 4},
		{[]int {1,2,2,3,4,2,3,6,7,2,4,1,8}, 1, 8},
		{[]int {1,2,2,3,4,2,3,6,7,2,4,1,8}, 13, 1},
		{[]int {123,231,432,10,123,123,123,10}, 5, 10},
	}

	for _, c := range cases {
		inNode := NodeSetupSmart(c.inList)
		got := inNode.KthToLast(c.inK)
		if got != c.want {
			t.Errorf("method KthToLast(%q) on list %q == %q, want %q", c.inK, inNode.ListString(), got, c.want)
		}
	}
}
