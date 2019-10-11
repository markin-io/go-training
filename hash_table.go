package main

import (
	"hash/fnv"
)

type HashNode struct {
	key   string
	value interface{}
	next  *HashNode
}

type HashFunction func(string, int) int

type HashTable struct {
	table        []*HashNode
	hashFunction HashFunction
}

const defaultArraySize int = 127

func defaultHashFunction(key string, size int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	uint8Hash := (uint8)(h.Sum32())
	normalizedHash := (float64)(uint8Hash) / 256

	return (int)(normalizedHash * (float64)(size))
}

func CreateHashTable(hashFunction HashFunction) *HashTable {
	if hashFunction == nil {
		hashFunction = defaultHashFunction
	}

	return &HashTable{
		table:        make([]*HashNode, defaultArraySize),
		hashFunction: defaultHashFunction,
	}
}

func (e HashTable) add(key string, value interface{}) {
	var index int
	// Calculate array index for specified key
	index = e.hashFunction(key, len(e.table))

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
	var index int
	// Calculate array index for specified key
	index = e.hashFunction(key, len(e.table))

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
