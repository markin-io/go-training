package datastructures

import "log"

// QueueNode for Queue data structure implementation
type QueueNode struct {
	value interface{}
	prev  *QueueNode
}

// Queue data structure implementation
type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (e *Queue) Push(value interface{}) {
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

func (e *Queue) Pop() *QueueNode {
	node := e.head
	e.head = e.head.prev
	return node
}

func (e *Queue) Print() {
	currentNode := e.head

	if currentNode == nil {
		return
	}

	var str string

	str += currentNode.value.(string)

	for currentNode.prev != nil {
		currentNode = currentNode.prev

		str += " " + currentNode.value.(string)
	}

	log.Println(str)
}

func (e *Queue) IsEmpty() bool {
	return e.head == nil
}
