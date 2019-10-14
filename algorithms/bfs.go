package algorithms

import (
	datastructures "github.com/markin-io/go-training/data-structures"
)

func bfs(graph *datastructures.HashTable, rootNode string, searchFor string) bool {
	searchQueue := &datastructures.Queue{}
	checkedNodes := datastructures.CreateHashTable(nil)

	addNodesToSearchQueue(searchQueue, []string{rootNode}, checkedNodes)

	for searchQueue.IsEmpty() != true {
		searchQueue.Print()
		name := searchQueue.Pop().Value.(string)
		if name == searchFor {
			return true
		}

		checkedNodes.Add(name, nil)

		addNodesToSearchQueue(searchQueue, graph.Get(name).Value.([]string), checkedNodes)
	}

	return false
}

func addNodesToSearchQueue(searchQueue *datastructures.Queue, nodes []string,
	checkedNodes *datastructures.HashTable) {
	for _, node := range nodes {
		if checkedNodes.Get(node) == nil {
			searchQueue.Push(node)
		}
	}
}
