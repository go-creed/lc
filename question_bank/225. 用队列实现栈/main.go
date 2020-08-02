package main

import "container/list"

type MyStack struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		list: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyStack) Push(x int) {
	this.list.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyStack) Pop() int {
	return this.list.Remove(this.list.Front()).(int)
}

/** Get the front element. */
func (this *MyStack) Peek() int {
	return this.list.Front().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}
