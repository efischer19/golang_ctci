package ctci_chapter3

import "golang_ctci/ch_2"

//a utility Stack, for use with chapter 3 questions
type Stack struct {
	top *ctci_chapter2.Node
}

func (stack *Stack) Pop() int {
	if stack.top != nil {
		ret := stack.top.Data()
		stack.top = stack.top.Next()
		return ret
	}
	panic("can't Pop() on an empty stack")
}

func (stack *Stack) Push(newData int) {
	newTop := ctci_chapter2.CreateNode(stack.top, newData)
	stack.top = &newTop
}

func (stack *Stack) Peek() int {
	return stack.top.Data()
}