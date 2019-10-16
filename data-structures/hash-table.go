package datastructures

import (
	"hash/fnv"
	"log"
)

type HashNodeValue interface{}

type HashNode struct {
	Key   string
	Value HashNodeValue
	Next  *HashNode
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

func (e *HashTable) Add(key string, value HashNodeValue) {
	var index int
	// Calculate array index for specified key
	index = e.hashFunction(key, len(e.table))

	// Check whether there's already node with the same key
	existingNode := e.Get(key)

	if existingNode == nil {
		// Create new node
		node := &HashNode{
			Key:   key,
			Value: value,
		}

		// Check whether nodes chain exists at specified index
		chainHead := e.table[index]

		// Assign new node as a chain head
		if chainHead != nil {
			node.Next = chainHead
		}

		e.table[index] = node
	} else {
		existingNode.Value = value
	}
}

func (e *HashTable) Get(key string) *HashNode {
	var index int
	// Calculate array index for specified key
	index = e.hashFunction(key, len(e.table))

	node := e.table[index]

	if node != nil {
		var foundNode *HashNode

		if node.Key == key {
			foundNode = node
		} else {
			currentNode := node
			for currentNode.Next != nil {
				currentNode = currentNode.Next
				if currentNode.Key == key {
					foundNode = currentNode
					break
				}

			}
		}

		return foundNode
	}

	return nil
}

func (e *HashTable) Keys() []string {
	keys := []string{}

	for _, headNode := range e.table {
		if headNode != nil {
			currentNode := &HashNode{
				Key:   "",
				Next:  headNode,
				Value: nil,
			}

			for currentNode.Next != nil {
				currentNode = currentNode.Next
				keys = append(keys, currentNode.Key)
			}
		}
	}

	return keys
}

func (e *HashTable) Print() {
	keys := e.Keys()

	for _, key := range keys {
		node := e.Get(key)
		log.Printf("Key %s, Value %v", key, node.Value)
	}
}
