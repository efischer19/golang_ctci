package ctci_chapter2

func (oldNode *Node) IsPalindrome() bool {
	//on review from the book, I could have done this better
	//I like their "slow-fast to find halfway point, make a stack up to there and continue to check front half vs back" approach
	//this solution of mine runs in O(n^2) time thanks to recurseReverse, and O(n) space thanks to constructing a copy of the list
	revOld := recurseReverse(*oldNode)
	for old, knew := oldNode, &revOld; old != nil && knew != nil; old, knew = old.next, knew.next {
		if old.data != knew.data {
			return false
		}
	}
	return true
}