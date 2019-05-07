package main

import "log"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

var head *Node
var tail *Node

func insert(value int) {
	node := &Node{
		value,
		nil,
		tail,
	}

	if head == nil {
		head = node
	} else {
		tail.next = node
	}

	tail = node
}

func runLinkedList() {
	insert(1)
	insert(2)
	insert(3)
	insert(4)

	currentNode := head

	log.Printf("Value %v", currentNode.value)
	for currentNode.next != nil {
		currentNode = currentNode.next
		log.Printf("Value %v", currentNode.value)
	}

	// Reversed order
	currentNode = tail
	log.Printf("Value rev %v", currentNode.value)
	for currentNode.prev != nil {
		currentNode = currentNode.prev
		log.Printf("Value rev %v", currentNode.value)
	}

}
