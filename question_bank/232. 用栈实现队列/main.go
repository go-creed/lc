package main

import (
	"container/list"
	"fmt"
)

func main() {
	constructor := Constructor()
	constructor.Push(1)
	constructor.Push(2)
	fmt.Println(constructor)
	peek := constructor.Peek()
	fmt.Println(peek)
	pop := constructor.Pop() // 返回 1
	fmt.Println(pop)
	empty := constructor.Empty()
	fmt.Println(empty)
}

type MyQueue struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		list: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.list.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.list.Remove(this.list.Front()).(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.list.Front().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.list.Len() == 0
}
