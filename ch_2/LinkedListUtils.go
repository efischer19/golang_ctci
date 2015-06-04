package ctci_chapter2

import (
	"bytes"
	"strconv"
)

//a utility linked list, for use with all of the chapter 2 questions
type Node struct {
	next *Node
	data int
}

//exporting accessors for use in later chapters
func (n Node) Next() *Node {
	return n.next
}

func (n Node) Data() int {
	return n.data
}

func CreateNode(next *Node, data int) Node{
	return Node{next, data}
}

//and one setter too
func (n *Node) SetNext(newNext *Node) {
	n.next = newNext
}

func (oldNode *Node) Append(newData int) {
	newNode :=  Node{nil, newData}
	newNode.next = oldNode.next
	oldNode.next = &newNode
}

func (oldNode *Node) AppendToTail(newData int) {
	newNode :=  Node{nil, newData}
	n := oldNode
	for ; n.next != nil; n = n.next {
	}
	n.next = &newNode
}

func NodeSetup(input []int) Node {
	if len(input) <= 0 {
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
	if len(input) <= 0 {
		panic("need to have elements to create linked list")
	}
	newNode := Node{nil, input[0]}
	curNode := &newNode
	for i:=1; i < len(input); i++ {
		curNode.Append(input[i])
		curNode = curNode.next
	}
	return newNode
}

func (oldNode Node) ListEquals(comp Node) bool {
	a, b := &oldNode, &comp
	for ; a != nil && b != nil; a, b = a.next, b.next{
		if a.data != b.data {
			return false
		}
	}
	return a == nil && b == nil
}

func (oldNode Node) ListString() string {
	var buffer bytes.Buffer
	buffer.WriteRune(' ')
	for i := &oldNode; i != nil; i = i.next {
		buffer.WriteString(strconv.Itoa(i.data))
		buffer.WriteRune(' ')
	}
	return buffer.String()
}