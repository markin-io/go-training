package algorithms

import (
	"log"
)

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	check := array[len(array)/2]
	less := []int{}
	more := []int{}
	checks := []int{}
	for _, num := range array {
		if num < check {
			less = append(less, num)
		} else if num > check {
			more = append(more, num)
		} else {
			checks = append(checks, check)
		}
	}

	less = quickSort(less)
	more = quickSort(more)

	less = append(less, checks...)
	return append(less, more...)
}

func main() {
	var data = ReadPipeInputInt()

	sorted := quickSort(data)

	log.Printf("Sorted output %v", sorted)
}
