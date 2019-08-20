package main

import (
	"log"
	"os"
	"strconv"
)

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

// remove
// search
// get for pos

func initialize() {
	var data = ReadPipeInput()

	for _, number := range data {
		insert(number)
	}
}

func getItemAt(index int) *Node {
	currentIndex := 0
	var foundNode *Node

	currentNode := head
	for currentNode.next != nil {
		if currentIndex == index {
			foundNode = currentNode
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return foundNode
}

func searchItem(value int) *Node {
	var foundNode *Node

	currentNode := head
	for currentNode.next != nil {
		if currentNode.value == value {
			foundNode = currentNode
			break
		}

		currentNode = currentNode.next
	}

	return foundNode
}

func removeItem(node *Node) {
	node.prev.next = node.next
	node = nil
}

func main() {
	initialize()

	// Searching
	var searchFor, _ = strconv.Atoi(string(os.Args[2]))
	item := searchItem(searchFor)
	if item != nil {
		log.Printf("Item with value %d found", searchFor)
	} else {
		log.Printf("Item with value %d not found", searchFor)
	}

	// Indexing
	var searchAt, _ = strconv.Atoi(string(os.Args[1]))
	item = getItemAt(searchAt)

	if item != nil {
		log.Printf("Item found at index %d has value %d", searchAt, item.value)
	} else {
		log.Printf("Item at index %d not found", searchAt)
	}

	// Removal
	removeItem(item)
	newItem := getItemAt(searchAt)

	// Check whether new links set up correctly
	log.Printf("Equal %v", item.prev.next == newItem)
	item = nil
}
