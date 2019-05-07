package main

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
