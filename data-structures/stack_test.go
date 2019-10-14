package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	stack := Stack{}

	for _, number := range numbers {
		stack.Push(number)
	}

	currentNode := stack.tail

	index := 0
	for currentNode.prev != nil {
		currentNode = stack.Pop()

		number := numbers[len(numbers)-index-1]
		if currentNode.value != number {
			t.Errorf("stack.pop() == %d, want %d", currentNode.value, number)
		}

		index++
	}
}
