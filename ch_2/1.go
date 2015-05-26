package ctci_chapter2

//this approach uses no extra space, but runs in O(n^2) time
func RemoveDupesNaive(n Node) Node {
	for cur := &n; cur != nil; cur = cur.next {
		for inner := cur; inner.next != nil; inner = inner.next {
			if cur.data == inner.next.data {
				inner.next = inner.next.next
			}
		}
	}
	return n
}

//this solution uses O(n) space, and runs in O(n) time
func RemoveDupes(n Node) Node {
	datas := make(map[int]bool)
	datas[n.data] = true
	for cur := &n; cur.next != nil; cur = cur.next {
		 if datas[cur.next.data] {
			cur.next = cur.next.next
		} else {
			datas[cur.next] = true
		}
	}
	return n
}
