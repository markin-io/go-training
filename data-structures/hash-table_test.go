package datastructures

import (
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {

	// Initialize and fill hash table
	hashTable := CreateHashTable(nil)

	for i := 0; i <= 1000; i++ {
		hashTable.Add(strconv.Itoa(i), i)
	}

	// Check get functionality
	for i := 0; i <= 1000; i++ {
		key := strconv.Itoa(i)
		node := hashTable.Get(key)

		if node.Value != i {
			t.Errorf("hashTable.get(%s).value == %d, wants %d", key, node.Value, i)
		}
	}
}

func TestReplacement(t *testing.T) {
	hashTable := CreateHashTable(nil)

	hashTable.Add("key", 1)
	hashTable.Add("key", 2)

	value := hashTable.Get("key").Value
	if value != 2 {
		t.Errorf("hashTable.get(%s).value == %d, wants %d", "key", value, 2)
	}
}
