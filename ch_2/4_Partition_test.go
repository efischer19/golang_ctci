package ctci_chapter2

import "testing"

func TestPartition(t *testing.T) {
	cases := []struct {
		in []int
		k int
	}{
		{[]int {10,1,9,2,8,3,7,4,6,5}, 5},
		{[]int {1,2,2,3,4,2,6}, 3},
	}

	for _, c := range cases {
		inNode := NodeSetupSmart(c.in)
		cur := &inNode
		gotNode := cur.Partition(c.k)
		if !gotNode.VerifyPartition(c.k) {
			t.Errorf("Partition(%q, %d) == %q", inNode.ListString(), c.k, gotNode.ListString())
		}
	}
}

func (oldNode *Node) VerifyPartition(k int) bool {
	cur := oldNode
	for ; cur != nil && cur.data < k; cur = cur.next {}
	for ; cur != nil; cur = cur.next {
		if cur.data < k {
			return false
		}
	}
	return true
}