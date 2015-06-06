package ctci_chapter3

//a stack which tracks min, with push(), pop() and min() all operating in O(1)
type StackWithMin struct {
	top *StackNode
}

type StackNode struct {
	next *StackNode
	data int
	min int
}

func (stack *StackWithMin) Pop() int {
	if stack.top != nil {
		ret := stack.top.data
		stack.top = stack.top.next
		return ret
	}
	panic("can't Pop() on an empty stack")
}

func (stack *StackWithMin) Push(newData int) {
	newMin := stack.top.min
	if newData < newMin {
		newMin = newData
	}
	newTop := StackNode{stack.top, newData, newMin}
	stack.top = &newTop
}

func (stack *StackWithMin) Min() int {
	return stack.top.min
}

func (stack *StackWithMin) Peek() int {
	return stack.top.data
}

//on review - this could have been more space efficient by allocating a second stack, and using that to track the mins as they change
//this solution is simple and readable though, so in the absence of other constraints I'm OK with it