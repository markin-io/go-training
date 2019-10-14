package algorithms

import (
	"testing"
)

func TestAscendingSort(t *testing.T) {
	numbers := []int{10, 8, 2, 1, 5, 3, 7, 6, 4, 9}

	sorted := quickSort(numbers)

	prev := sorted[0]

	for _, value := range sorted {
		if prev > value {
			t.Errorf("quickSort() error, expects %d >= %d", value, prev)
		}
		prev = value
	}
}
