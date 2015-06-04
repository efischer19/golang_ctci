package ctci_chapter2

func AddNumbersOnesFirst(a Node, b Node) Node {
	carry := 0
	ret := Node{nil, -1}
	aPtr, bPtr := &a, &b
	for ; aPtr != nil && bPtr != nil; aPtr, bPtr = aPtr.next, bPtr.next {
		sum := aPtr.data + bPtr.data + carry
		if sum > 9 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}
		if ret.data == -1 {
			ret.data = sum
		} else {
			ret.AppendToTail(sum)
		}
	}
	
	//only 1 of these next 2 loops may run, do to condition of first for loop
	for ; aPtr != nil; aPtr = aPtr.next {
		sum := aPtr.data + carry
		if sum > 9 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}
		ret.AppendToTail(sum)
	}
	
	for ; bPtr != nil; bPtr = bPtr.next {
		sum := bPtr.data + carry
		if sum > 9 {
			sum %= 10
			carry = 1
		} else {
			carry = 0
		}
		ret.AppendToTail(sum)
	}
	
	if carry != 0 {
		ret.AppendToTail(carry)
	}
	
	return ret
}

func AddNumbersOnesLast(a Node, b Node) Node {
	//this is kind of cheating, but if I were given the problems in this order, this is how I would solve them
	//since I already have a working AddOnesFirst, simply reversing the list is easier/more eficient use of my dev time than trying to sort out all the details of OnesLast
	//but if I were to do that, it would involve some way of passing the 2 lists into goroutines, and having those pass back the partial sums in ones first order
	aRev := recurseReverse(a)
	bRev := recurseReverse(b)
	return recurseReverse(AddNumbersOnesFirst(aRev, bRev))
}

func recurseReverse(n Node) Node {
	if n.next == nil {
		return n
	}
	ret := recurseReverse(*(n.next))
	
	//this causes recurseReverse to run in O(n^2), I'm banking on n being small enough to not matter much (# of digits)
	ret.AppendToTail(n.data)
	return ret
}