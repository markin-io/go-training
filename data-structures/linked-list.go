package datastructures

type LinkedListNode struct {
	value int
	next  *LinkedListNode
	prev  *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (e *LinkedList) insert(value int) {
	node := &LinkedListNode{
		value,
		nil,
		e.tail,
	}

	if e.head == nil {
		e.head = node
	} else {
		e.tail.next = node
	}

	e.tail = node
}

func (e *LinkedList) getItemAt(index int) *LinkedListNode {
	currentIndex := 0
	var foundNode *LinkedListNode

	currentNode := e.head
	for currentNode.next != nil {
		if currentIndex == index {
			foundNode = currentNode
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return foundNode
}

func (e *LinkedList) searchItem(value int) *LinkedListNode {
	var foundNode *LinkedListNode

	currentNode := e.head
	for currentNode.next != nil {
		if currentNode.value == value {
			foundNode = currentNode
			break
		}

		currentNode = currentNode.next
	}

	return foundNode
}

func (e *LinkedList) removeItem(node *LinkedListNode) {
	node.prev.next = node.next
	node = nil
}
