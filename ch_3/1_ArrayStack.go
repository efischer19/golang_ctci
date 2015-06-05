package ctci_chapter3

type ArrayStack struct {
	data [1002]int	//chosen to be a reasonably large divisible by 3 size
	topDex [3]int
	lenData, numSubs, subLen int
}

func (aStack *ArrayStack) Pop(subStack int) int {
	aStack.BasicChecks(subStack)
	
	min := aStack.subLen*subStack
	if aStack.topDex[subStack] < min {
		panic("stack underflow")
	}
	
	aStack.topDex[subStack]--
	return aStack.data[aStack.topDex[subStack+1]]
}

func (aStack *ArrayStack) Push(subStack, newData int) {
	aStack.BasicChecks(subStack)
	
	max := (aStack.subLen*(subStack+1))
	if aStack.topDex[subStack] >= max {
		panic("stack overflow")
	}
	
	aStack.topDex[subStack]++
	aStack.data[aStack.topDex[subStack]] = newData
}

func (aStack *ArrayStack) BasicChecks(i int) {
	if aStack.subLen == 0 {
		aStack.Setup()
	}
	if i > aStack.numSubs {
		panic("Index out of range, there are 3 substacks per ArrayStack")
	}
}

func (aStack *ArrayStack) Setup() {
	aStack.lenData, aStack.numSubs = len(aStack.data), len(aStack.topDex)
	aStack.subLen = aStack.lenData/aStack.numSubs
}