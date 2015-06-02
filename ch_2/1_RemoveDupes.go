package ctci_chapter2

//this approach uses no extra space, but runs in O(n^2) time
func (n *Node) RemoveDupesNaive() {
	for cur := n; cur != nil; cur = cur.next {
		for inner := cur; inner != nil && inner.next != nil; inner = inner.next {
			if cur.data == inner.next.data {
				inner.next = inner.next.next
			}
		}
	}
}

//this solution uses O(n) space, and runs in O(n) time
func (n *Node) RemoveDupes() {
	datas := make(map[int]bool)
	datas[n.data] = true
	for cur := n; cur.next != nil; {
		 if datas[cur.next.data] {
			cur.next = cur.next.next
		} else {
			datas[cur.next.data] = true
			cur = cur.next
		}
	}
}
