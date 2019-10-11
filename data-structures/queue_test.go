package datastructures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	queue := Queue{}

	for _, number := range numbers {
		queue.push(number)
	}

	currentNode := queue.head

	index := 0
	for currentNode.prev != nil {
		currentNode = queue.pop()

		number := numbers[index]
		if currentNode.value != number {
			t.Errorf("queue.pop() == %d, want %d", currentNode.value, number)
		}

		index++
	}
}
