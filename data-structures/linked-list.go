package datastructures

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
	Prev  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

func (e *LinkedList) Insert(value int) {
	node := &LinkedListNode{
		value,
		nil,
		e.Tail,
	}

	if e.Head == nil {
		e.Head = node
	} else {
		e.Tail.Next = node
	}

	e.Tail = node
}

func (e *LinkedList) GetItemAt(index int) *LinkedListNode {
	currentIndex := 0
	var foundNode *LinkedListNode

	currentNode := e.Head
	for currentNode.Next != nil {
		if currentIndex == index {
			foundNode = currentNode
		}

		currentNode = currentNode.Next
		currentIndex++
	}

	return foundNode
}

func (e *LinkedList) SearchItem(value int) *LinkedListNode {
	var foundNode *LinkedListNode

	currentNode := e.Head
	for currentNode.Next != nil {
		if currentNode.Value == value {
			foundNode = currentNode
			break
		}

		currentNode = currentNode.Next
	}

	return foundNode
}

func (e *LinkedList) RemoveItem(node *LinkedListNode) {
	node.Prev.Next = node.Next
	node = nil
}
