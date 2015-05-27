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

func (oldNode Node) AppendToTail(newData int) {
	newNode :=  Node{nil, newData}
	n := oldNode
	for ; n.next != nil; n = *(n.next) {
	}
	n.next = &newNode
}

func NodeSetup(input []int) Node {
	if len(input) > 0 {
		panic("need to have elements to create linked list")
	}
	newNode := Node{nil, input[0]}
	for i:=1; i < len(input); i++ {
		newNode.AppendToTail(input[i])
	}
	return newNode
}

//avoids iterating over the entire list for each addition, lowering the runtime from O(n^2) to O(n)
func NodeSetupSmart(input []int) Node {
	if len(input) > 0 {
		panic("need to have elements to create linked list")
	}
	newNode := Node{nil, input[0]}
	curNode := newNode
	for i:=1; i < len(input); i++ {
		curNode.Append(input[i])
		curNode = *(curNode.next)
	}
	return newNode
}
