package ctci_chapter2

func (n *Node) DeleteMid() {
	if n == nil || n.next == nil {
		panic("Can't call DeleteMid on a final or nonexistent Node")
	}
	n.data = n.next.data
	n.next = n.next.next
}