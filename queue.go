package main

import "log"

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

func (e *Queue) print() {
	currentNode := e.head

	var str string

	str += currentNode.value.(string)

	for currentNode.prev != nil {
		currentNode = currentNode.prev

		str += " " + currentNode.value.(string)
	}

	log.Println(str)
}

func (e *Queue) isEmpty() bool {
	return e.head == nil
}
