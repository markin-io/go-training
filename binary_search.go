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

func main() {
	var data = ReadPipeInput()

	var toFind, _ = strconv.Atoi(string(os.Args[1]))
	log.Printf("Search for %v", toFind)

	index := binarySearch(data, toFind)
	log.Printf("Index %d", index)
}
