package main

import (
	"hash/fnv"
	"log"
)

var tableSize int

type HashNode struct {
	key   string
	value int
	next  *HashNode
}

type HashTable struct {
	table []*HashNode
}

func (e HashTable) add(key string, value int) {
	// Calculate array index for specified key
	index := hashFunc(key)

	// Check whether there's already node with the same key
	existingNode := e.get(key)

	if existingNode == nil {
		// Create new node
		node := &HashNode{
			key:   key,
			value: value,
		}

		// Check whether nodes chain exists at specified index
		chainHead := e.table[index]

		// Assign new node as a chain head
		if chainHead != nil {
			node.next = chainHead
		}

		e.table[index] = node
	} else {
		existingNode.value = value
	}
}

func (e HashTable) get(key string) *HashNode {
	index := hashFunc(key)

	node := e.table[index]

	if node != nil {
		var foundNode *HashNode

		if node.key == key {
			foundNode = node
		} else {
			currentNode := node
			for currentNode.next != nil {
				currentNode = currentNode.next
				if currentNode.key == key {
					foundNode = currentNode
					break
				}

			}
		}

		return foundNode
	}

	return nil
}

func createHashTable(size int) *HashTable {
	return &HashTable{
		table: make([]*HashNode, size),
	}
}

func hashFunc(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	uint8Hash := (uint8)(h.Sum32())
	normalizedHash := (float64)(uint8Hash) / 256

	return (int)(normalizedHash * (float64)(tableSize))
}

func main() {
	// Read array of random strings
	var data = ReadPipeInputString()

	tableSize = len(data)

	// Initialize and fill hash table
	hashTable := createHashTable(tableSize)
	for index, key := range data {
		hashTable.add(key, index)
	}

	// Check get functionality
	for index, key := range data {
		node := hashTable.get(key)
		if node.value != index {
			log.Printf("Hashtable missed value!")
		}
	}

	// Check whether hashtable correctly replace values with the same key
	hashTable.add("testKey", 1)
	hashTable.add("testKey", 2)

	if hashTable.get("testKey").value != 2 {
		log.Printf("Incorrect replace!")
	}
}
