package main

import (
	"log"
	"os"
	"strconv"
)

// implement avg and worst cases https://en.wikipedia.org/wiki/Best,_worst_and_average_case#Data_structures

func binarySearch(array []int, value int) int {
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

func binarySearchRecursive(array []int, value int) int {
	if len(array) == 0 || array[0] > value || array[len(array)-1] < value {
		return -1
	}

	low := 0
	high := len(array) - 1

	mid := (low + high) / 2

	guess := array[mid]

	if guess == value {
		return mid
	} else if guess > value {
		return binarySearchRecursive(append([]int{}, array[0:mid]...), value)
	} else if guess < value {
		return mid + 1 + binarySearchRecursive(append([]int{}, array[mid+1:]...), value)
	}

	return -1
}

func main() {
	var data = ReadPipeInput()

	var toFind, _ = strconv.Atoi(string(os.Args[1]))
	log.Printf("Search for %d", toFind)

	searchIndex := binarySearch(data, toFind)
	recursiveSearchIndex := binarySearchRecursive(data, toFind)

	log.Printf("Index is %d and %d", searchIndex, recursiveSearchIndex)
}
