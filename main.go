package main

// func runQueue() {
// 	log.Println("\nRun queue")
// 	queue := Queue{}

// 	queue.push(1)
// 	queue.push("string")
// 	queue.push(3)
// 	queue.push(4)

// 	currentNode := queue.pop()

// 	log.Printf("Value %v", currentNode.value)
// 	for currentNode.prev != nil {
// 		currentNode = queue.pop()

// 		log.Printf("Value %v", currentNode.value)
// 	}
// }

// func runStack() {
// 	log.Println("\nRun stack")
// 	stack := Stack{}
// 	stack.push(1)
// 	stack.push(2)
// 	stack.push(3)
// 	stack.push(4)

// 	currentNode := stack.pop()

// 	log.Printf("Value %v", currentNode.value)
// 	for currentNode.prev != nil {
// 		currentNode = stack.pop()
// 		log.Printf("Value %v", currentNode.value)
// 	}

// }

// func runHashTable() {
// 	log.Println("\nRun hash table")
// 	// Read array of random strings
// 	var data = ReadPipeInputString()

// 	// Initialize and fill hash table
// 	hashTable := &HashTable{
// 		table: make([]*HashNode, len(data)),
// 	}

// 	for index, key := range data {
// 		hashTable.add(key, index)
// 	}

// 	// Check get functionality
// 	for index, key := range data {
// 		node := hashTable.get(key)

// 		if node.value != index {
// 			log.Printf("Hashtable missed value!")
// 		}
// 	}

// 	// Check whether hashtable correctly replace values with the same key
// 	hashTable.add("testKey", 1)
// 	hashTable.add("testKey", 2)

// 	if hashTable.get("testKey").value != 2 {
// 		log.Printf("Incorrect replace!")
// 	}
// }

func main() {
	// runQueue()
	// runStack()
	// runHashTable()
	runBfs()
}
