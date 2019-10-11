package main

import "log"

func runBfs() {
	graph := &HashTable{
		table: make([]*HashNode, 8),
	}

	// Build users graph
	graph.add("you", []string{"alice", "bob", "claire"})
	graph.add("bob", []string{"anuj", "peggy"})
	graph.add("alice", []string{"peggy"})
	graph.add("claire", []string{"tom", "jonny"})
	graph.add("anuj", []string{})
	graph.add("peggy", []string{})
	graph.add("tom", []string{})
	graph.add("jonny", []string{})

	// Find tom
	searchQueue := &Queue{}
	addNodesToSearchQueue(searchQueue, graph.get("you").value.([]string))
	for searchQueue.isEmpty() != true {
		searchQueue.print()
		name := searchQueue.pop().value.(string)

		if name == "tom" {
			log.Println("Found tom")
			return
		}

		addNodesToSearchQueue(searchQueue, graph.get(name).value.([]string))
	}
}

func addNodesToSearchQueue(queue *Queue, nodes []string) {
	for _, node := range nodes {
		// log.Println(node)
		queue.push(node)
	}
}
