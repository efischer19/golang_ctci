package ctci_chapter3

//a set of stacks, as outlined in problem 3.3
type SetOfStacks struct {
	top *StackWithSize
}

type StackWithSize struct {
	innerStack *Stack
	next *StackWithSize
	size int
}

//intentionally using a very low value for testing
var SetOfStacksCapacity int = 5

//StackWithSize.Pop, return stack.Pop() and decrement size
func (stack *StackWithSize) Pop() int {
	stack.size--
	return stack.innerStack.Pop()
}

//SetOfStacks.Pop(), call StackWithSize.Pop() and update top if size == 0 after
func (stack *SetOfStacks) Pop() int {
	if stack.top != nil {
		ret := stack.top.innerStack.Pop()
		if stack.top.size == 0 {
			stack.top = stack.top.next
		}
		return ret
	}
	panic("can't Pop() on an empty stack")
}

//same with Push for both data structures
func (stack *StackWithSize) Push(newData int) {
	stack.innerStack.Push(newData)
	stack.size++
}

func (stack *SetOfStacks) Push(newData int) {
	if stack.top == nil || stack.top.size == SetOfStacksCapacity {
		newTop := StackWithSize{&Stack{nil}, stack.top, 1}
		stack.top = &newTop
	}
	stack.top.Push(newData)
}