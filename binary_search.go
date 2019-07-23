package main

import "log"

// implement files i/o and tests
// implement avg and worst cases https://en.wikipedia.org/wiki/Best,_worst_and_average_case#Data_structures

func binarySearch(array [11]int, value int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		guess := array[mid]

		if guess == value {
			return mid
		} else if guess > value {
			high = mid - 1
		} else if guess < value {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	var array = [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := binarySearch(array, 9)
	log.Printf("Index %d", index)
}
