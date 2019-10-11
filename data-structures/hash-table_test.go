package datastructures

import (
	"testing"
)

func TestGet(t *testing.T) {
	strings := []string{"one", "two", "three", "four", "five"}

	// Initialize and fill hash table
	hashTable := CreateHashTable(nil)

	for index, value := range strings {
		hashTable.add(value, index)
	}

	// Check get functionality
	for index, key := range strings {
		node := hashTable.get(key)

		if node.value != index {
			t.Errorf("hashTable.get(%s).value == %d, wants %d", key, node.value, index)
		}
	}
}

func TestReplacement(t *testing.T) {
	hashTable := CreateHashTable(nil)

	hashTable.add("key", 1)
	hashTable.add("key", 2)

	value := hashTable.get("key").value
	if value != 2 {
		t.Errorf("hashTable.get(%s).value == %d, wants %d", "key", value, 2)
	}
}
