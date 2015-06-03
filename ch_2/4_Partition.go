package ctci_chapter2

func (n *Node) Partition(k int) Node {
	var low, high, uninit Node
	for ; n != nil; n = n.next {
		if n.data < k {
			if low == uninit {
				low = Node{nil, n.data}
			} else {
				low.Append(n.data)
			}
		} else {
			if high == uninit {
				high = Node{nil, n.data}
			} else {
				high.Append(n.data)
			}
		}
	}
	
	cur := &low
	for ; cur.next != nil; cur = cur.next {}
	cur.next = &high
	return low
}