package main

import "container/list"

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	stack *list.List
	it    *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	stack := list.New()
	iterator := &PeekingIterator{
		stack: stack,
		it:    iter,
	}
	if iter.hasNext() {
		stack.PushFront(iter.next())
	}

	return iterator
}

func (this *PeekingIterator) hasNext() bool {
	if this.stack.Len() > 0 {
		return true
	}
	//if this.stack.Len() == 0 || !this.it.hasNext() {
	//	return false
	//}
	return this.it.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.it.hasNext() {
		this.stack.PushBack(this.it.next())
	}
	remove := this.stack.Remove(this.stack.Front()).(int)
	return remove
}

func (this *PeekingIterator) peek() int {
	return this.stack.Front().Value.(int)
}
