package datastructures

type StackNode struct {
	value interface{}
	prev  *StackNode
}

type Stack struct {
	tail *StackNode
}

func (e *Stack) push(value interface{}) {
	node := &StackNode{
		value,
		e.tail,
	}

	e.tail = node
}

func (e *Stack) pop() *StackNode {
	node := e.tail
	e.tail = e.tail.prev
	return node
}
