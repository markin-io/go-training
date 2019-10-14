package datastructures

type StackNode struct {
	Value interface{}
	Prev  *StackNode
}

type Stack struct {
	Tail *StackNode
}

func (e *Stack) Push(value interface{}) {
	node := &StackNode{
		value,
		e.Tail,
	}

	e.Tail = node
}

func (e *Stack) Pop() *StackNode {
	node := e.Tail
	e.Tail = e.Tail.Prev
	return node
}
