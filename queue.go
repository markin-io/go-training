package main

type QueueNode struct {
	value interface{}
	prev  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (e *Queue) push(value interface{}) {
	node := &QueueNode{
		value,
		nil,
	}

	if e.head == nil {
		e.head = node
	}

	if e.tail != nil {
		e.tail.prev = node
	}

	e.tail = node
}

func (e *Queue) pop() *QueueNode {
	node := e.head
	e.head = e.head.prev
	return node
}