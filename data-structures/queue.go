package datastructures

import "log"

// QueueNode for Queue data structure implementation
type QueueNode struct {
	Value interface{}
	Prev  *QueueNode
}

// Queue data structure implementation
type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func (e *Queue) Push(value interface{}) {
	node := &QueueNode{
		value,
		nil,
	}

	if e.Head == nil {
		e.Head = node
	}

	if e.Tail != nil {
		e.Tail.Prev = node
	}

	e.Tail = node
}

func (e *Queue) Pop() *QueueNode {
	node := e.Head
	e.Head = e.Head.Prev
	return node
}

func (e *Queue) Print() {
	currentNode := e.Head

	if currentNode == nil {
		return
	}

	var str string

	str += currentNode.Value.(string)

	for currentNode.Prev != nil {
		currentNode = currentNode.Prev

		str += " " + currentNode.Value.(string)
	}

	log.Println(str)
}

func (e *Queue) IsEmpty() bool {
	return e.Head == nil
}
