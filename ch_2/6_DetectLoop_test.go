package ctci_chapter2

import "testing"

func TestDetectLoop(t *testing.T) {
	cases := []struct {
		in []int
		loopPoint int	//zero-indexed, which node tail points to. This is also the 'data' of the desired return Node
	}{
		{[]int {1,2,3,4,5}, 1},
		{[]int {1,2,3,4,5}, 2},
		{[]int {0,9,8,7,6,5,4,3,2,1,10,345,35,123,765,234,765}, 11},
	}

	for _, c := range cases {
		inNode := NodeSetupSmart(c.in)
		priorString := inNode.ListString()	//so ListString doesn't explode later
		inNode.CreateLoop(c.loopPoint)
		got := DetectLoop(inNode)
		if got.data != c.in[c.loopPoint] {
			t.Errorf("DetectLoop(%q %d) == %d, want %d", priorString, c.in[c.loopPoint], got.data, c.in[c.loopPoint])
		}
	}
}

func (oldNode *Node) CreateLoop (loopPoint int) {
	var loopStart, cur *Node = nil, oldNode
	for i := 0; cur.next != nil; cur, i = cur.next, i+1 {
		if i == loopPoint {
			loopStart = cur
		}
	}
	
	if loopStart == nil {
		panic("Reached end of list before loopPoint")
	}
	
	cur.next = loopStart
}