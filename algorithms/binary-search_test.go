package algorithms

import (
	"testing"
)

func TestSearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	toSearch := 5
	index := 4
	foundIndex := binarySearch(numbers, toSearch)

	if foundIndex != index {
		t.Errorf("binarySearch() index == %d, wants %d", foundIndex, index)
	}
}

func TestRecursiveSearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	toSearch := 5
	index := 4
	foundIndex := binarySearchRecursive(numbers, toSearch)

	if foundIndex != index {
		t.Errorf("binarySearch() index == %d, wants %d", foundIndex, index)
	}
}
