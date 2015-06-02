package ctci_chapter2

import "testing"

func TestRemoveDupes(t *testing.T) {
	cases := []struct {
		in []int
		want []int
	}{
		{[]int {1,2,2,3,4,2,3,6,7,2,4,1,8}, []int {1,2,3,4,6,7,8}},
		{[]int {123,231,432,10,123,123,123,10}, []int {123,231,432,10}},
	}

	for _, c := range cases {
		inNode, wantNode := NodeSetupSmart(c.in), NodeSetupSmart(c.want)
		priorString := inNode.ListString()
		inNode.RemoveDupesNaive()
		if !inNode.ListEquals(wantNode) {
			t.Errorf("RemoveDupesNaive(%q) == %q, want %q", priorString, inNode.ListString(), wantNode.ListString())
		}
		
		inNode = NodeSetupSmart(c.in)
		priorString = inNode.ListString()
		inNode.RemoveDupes()
		if !inNode.ListEquals(wantNode) {
			t.Errorf("RemoveDupes(%q) == %q, want %q", priorString, inNode.ListString(), wantNode.ListString())
		}
	}
}
