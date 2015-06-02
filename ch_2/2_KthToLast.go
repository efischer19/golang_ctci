package ctci_chapter2

func (n *Node) KthToLast(k int) int {	
	//place 'cur' k places in front of n
	cur := n
	for i := 1; cur != nil && i < k; cur, i = cur.next, i+1 {}
	
	//panic if needed
	if cur == nil {
		panic("List shorted than provided k")
	}
	
	//now just iterate to the end
	for ; cur.next != nil; cur, n = cur.next, n.next {}
	
	return n.data
}