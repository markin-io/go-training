package algorithms

import (
	datastructures "github.com/markin-io/go-training/data-structures"
)

func bfsFindPath(graph *datastructures.HashTable, src string, dest string) ([]string, int) {
	distances := datastructures.CreateHashTable(nil)
	predcessors := datastructures.CreateHashTable(nil)
	visitedNodes := datastructures.CreateHashTable(nil)
	searchQueue := datastructures.Queue{}

	searchQueue.Push(src)
	visitedNodes.Add(src, true)
	distances.Add(src, 0)

	for searchQueue.IsEmpty() != true {
		currentNode := searchQueue.Pop().Value.(string)

		if currentNode == dest {
			break
		}

		childNodes := graph.Get(currentNode).Value.([]string)

		for _, childNode := range childNodes {
			if visitedNodes.Get(childNode) == nil {
				searchQueue.Push(childNode)
				visitedNodes.Add(childNode, true)

				distances.Add(childNode, distances.Get(currentNode).Value.(int)+1)
				predcessors.Add(childNode, currentNode)
			}
		}
	}

	shortestDistance := distances.Get(dest).Value.(int)

	path := []string{}
	currentNode := dest
	path = append(path, dest)
	for currentNode != src {
		currentNode = predcessors.Get(currentNode).Value.(string)
		path = append(path, currentNode)
	}

	return path, shortestDistance
}

func bfsIsPathExists(graph *datastructures.HashTable, src string, dest string) bool {
	result := false

	searchQueue := datastructures.Queue{}
	visitedNodes := datastructures.CreateHashTable(nil)

	searchQueue.Push(src)
	visitedNodes.Add(src, true)

	for searchQueue.IsEmpty() != true {
		currentNode := searchQueue.Pop().Value.(string)

		if currentNode == dest {
			result = true
			break
		}

		childNodes := graph.Get(currentNode).Value.([]string)

		for _, childNode := range childNodes {
			if visitedNodes.Get(childNode) == nil {
				searchQueue.Push(childNode)
				visitedNodes.Add(childNode, true)
			}
		}
	}

	return result
}

func addNodesToSearchQueue(searchQueue *datastructures.Queue, nodes []string,
	checkedNodes *datastructures.HashTable) {
	for _, node := range nodes {
		if checkedNodes.Get(node) == nil {
			searchQueue.Push(node)
		}
	}
}
