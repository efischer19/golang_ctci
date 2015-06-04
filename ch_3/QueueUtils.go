package ctci_chapter3

import "golang_ctci/ch_2"

//a utility Queue, for use with chapter 3 questions
type Queue struct {
	first, last *ctci_chapter2.Node
}

func (q *Queue) Enqueue(newData int) {
	knew := ctci_chapter2.CreateNode(nil, newData)
	if q.first == nil {
		q.last = &knew
		q.first = q.last
	} else {
		q.last.SetNext(&knew)
		q.last = q.last.Next()
	}
}

func (q *Queue) Dequeue() int {
	if q.first != nil {
		ret := q.first.Data()
		q.first = q.first.Next()
		if q.first == nil { q.last = nil }	//empty Queue
		return ret
	}
	panic("Queue is empty")
}