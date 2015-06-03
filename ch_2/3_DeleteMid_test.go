package ctci_chapter2

import "testing"

func TestDeleteMid(t *testing.T) {
	cases := []struct {
		in []int
		inDel int	//number of cur->next iterations to use before calling DeleteMid()
		want []int
	}{
		{[]int {1,2,2,3,4}, 3, []int {1,2,2,4}},
		{[]int {1,2,2,3,4}, 0, []int {2,2,3,4}},
	}

	for _, c := range cases {
		inNode, wantNode := NodeSetupSmart(c.in), NodeSetupSmart(c.want)
		priorString := inNode.ListString()
		
		cur := &inNode
		//deliberately excluding a check for cur == nil, that's a panic that test cases should avoid hitting on their own
		for i := 0; i < c.inDel; cur, i = cur.next, i+1 {}
		cur.DeleteMid()
		if !inNode.ListEquals(wantNode) {
			t.Errorf("DeleteMid(%q, %d) == %q, want %q", priorString, c.inDel, inNode.ListString(), wantNode.ListString())
		}
	}
}