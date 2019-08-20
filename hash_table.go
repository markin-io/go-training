package main

import (
	"hash/fnv"
	"math"
)

var tableSize int

func hashFunc(key string) float64 {
	h := fnv.New32a()
	h.Write([]byte(key))
	uint8Hash := (uint8)(h.Sum32())
	normalizedHash := (float64)(uint8Hash) / 255

	return math.Round(normalizedHash * (float64)(tableSize))
}

func main() {
	tableSize = 100
	// tableSize, _ = strconv.Atoi(string(os.Args[1]))
}
