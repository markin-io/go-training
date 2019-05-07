package main

import "log"

type Node struct {
	value int
	prev  *Node
}

var tail *Node

func push(value int) {
	node := &Node{
		value,
		tail,
	}

	tail = node
}

func pop() *Node {
	node := tail
	tail = tail.prev
	return node
}

func main() {
	push(1)
	push(2)
	push(3)
	push(4)

	currentNode := pop()

	log.Printf("Value %v", currentNode.value)
	for currentNode.prev != nil {
		currentNode = pop()
		log.Printf("Value %v", currentNode.value)
	}
}
