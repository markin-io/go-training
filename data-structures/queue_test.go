package datastructures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	queue := Queue{}

	for _, number := range numbers {
		queue.Push(number)
	}

	currentNode := queue.Head

	index := 0
	for currentNode.Prev != nil {
		currentNode = queue.Pop()

		number := numbers[index]
		if currentNode.Value != number {
			t.Errorf("queue.pop() == %d, want %d", currentNode.Value, number)
		}

		index++
	}
}
