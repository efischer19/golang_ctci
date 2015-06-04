package ctci_chapter2

import "testing"

func TestAddNumbers(t *testing.T) {
	cases := []struct {
		inA []int
		inB []int
		want []int
	}{
		{[]int {7,1,6}, []int {5,9,2}, []int {2,1,9}},
		{[]int {9,1,1}, []int {1,9,9}, []int {0,1,1,1}},
		{[]int {5,4}, []int {5,5,9,1}, []int {0,0,0,2}},
	}

	for _, c := range cases {
		inNodeA, inNodeB, wantNode := NodeSetupSmart(c.inA), NodeSetupSmart(c.inB), NodeSetupSmart(c.want)
		got1 := AddNumbersOnesFirst(inNodeA, inNodeB)
		if !got1.ListEquals(wantNode) {
			t.Errorf("AddNumbersOnesFirst(%q, %q) == %q, want %q", inNodeA.ListString(), inNodeB.ListString(), got1.ListString(), wantNode.ListString())
		}
		
		Reverse(&c.inA)
		Reverse(&c.inB)
		Reverse(&c.want)
		inNodeA, inNodeB, wantNode = NodeSetupSmart(c.inA), NodeSetupSmart(c.inB), NodeSetupSmart(c.want)
		got2 := AddNumbersOnesLast(inNodeA, inNodeB)
		if !got2.ListEquals(wantNode) {
			t.Errorf("AddNumbersOnesLast(%q, %q) == %q, want %q", inNodeA.ListString(), inNodeB.ListString(), got2.ListString(), wantNode.ListString())
		}
	}
}

func Reverse(rPtr *[]int) {
	r := *rPtr
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}