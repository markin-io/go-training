package main

func runBfs() {
	graph := &HashTable{
		table: make([]*HashNode, 8),
	}

	// Build social graph
	graph.add("you", []string{"alice", "bob", "claire"})
	graph.add("bob", []string{"anuj", "peggy"})
	graph.add("alice", []string{"peggy"})
	graph.add("claire", []string{"tom", "jonny"})
	graph.add("anuj", []string{})
	graph.add("peggy", []string{})
	graph.add("tom", []string{})
	graph.add("jonny", []string{})

	searchQueue := &Queue{}

	addNodesToSearchQueue(searchQueue, graph.get("you").value.([]string))

	// searchQueue.print()
	// log.Printf("Quee %v", searchQueue)
}

func addNodesToSearchQueue(queue *Queue, nodes []string) {
	for _, node := range nodes {
		// log.Println(node)
		queue.push(node)
	}
}
