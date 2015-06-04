package ctci_chapter2

/*
The approach here: I use the standard fast-slow runner technique to detect a looped list
To find the beginning, I make use of the following two facts I believe to be true:
	-if fast and slow meet, the node they meet at is in the loop
	-if I remove* the "beginning" of the loop, there will no longer be a loop
		-where 'remove' means 'skip over it in the beginning of the loop, and point the end-of-loop instance to nil'
*/

func DetectLoop(n Node) *Node {
	if n.next == nil {
		panic("Can't find loops in list of length 1")
	}
	var slow, fast, prevS  *Node = &n, n.next, nil
	for ; slow != fast && fast != nil && fast.next != nil; prevS, slow, fast = slow, slow.next, fast.next.next {}
	if fast == nil || fast.next == nil {
		//this means we reached the end, and there are no loops
		return nil
	} else {
		//remove slow node, if this "fixes" the loop we've found our start point
		prevS.next = slow.next
		fast.next = nil
		possibleRet := DetectLoop(n)
		if possibleRet == nil {
			return slow
		} else {
			return possibleRet
		}
	}
}