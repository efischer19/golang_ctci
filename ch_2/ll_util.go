package ctci_chapter2

//a utility linked list, for use with all of the chapter 2 questions
type Node struct {
	next *Node
	data int
}

func (oldNode Node) Append(newData int) {
	newNode :=  Node{nil, newData}
	newNode.next = oldNode.next
	oldNode.next = &newNode
}
