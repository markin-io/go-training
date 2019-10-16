package algorithms

import (
	"math"

	datastructures "github.com/markin-io/go-training/data-structures"
)

func dijkstra(graph *datastructures.HashTable, src string, dest string) (float64, []string) {
	costs := datastructures.CreateHashTable(nil)
	predcessors := datastructures.CreateHashTable(nil)
	visitedNodes := datastructures.CreateHashTable(nil)

	// Fill initial costs of direct neighbors with their actual values.
	sourceNodeNeighbors := graph.Get(src).Value.(*datastructures.HashTable)
	for _, node := range sourceNodeNeighbors.Keys() {
		costs.Add(node, sourceNodeNeighbors.Get(node).Value)
	}

	// Non-direct neighbors getting +Inf
	for _, key := range graph.Keys() {
		if costs.Get(key) == nil && key != src {
			costs.Add(key, math.Inf(1))
		}
	}

	nodeKey := findLowestCostNode(costs, visitedNodes)
	for nodeKey != "" {
		cost := costs.Get(nodeKey).Value.(float64)
		neighborsNode := graph.Get(nodeKey).Value

		if neighborsNode != nil {
			neighbors := neighborsNode.(*datastructures.HashTable)
			for _, neighborNodeKey := range neighbors.Keys() {
				newCost := cost + neighbors.Get(neighborNodeKey).Value.(float64)

				if costs.Get(neighborNodeKey).Value.(float64) > newCost {
					costs.Add(neighborNodeKey, newCost)
					predcessors.Add(neighborNodeKey, nodeKey)
				}
			}
		}

		visitedNodes.Add(nodeKey, true)
		nodeKey = findLowestCostNode(costs, visitedNodes)
	}

	cost := costs.Get(dest).Value.(float64)
	path := getPath(predcessors, dest, src)

	return cost, path
}

func findLowestCostNode(costs *datastructures.HashTable, visitedNodes *datastructures.HashTable) string {
	keys := costs.Keys()

	lowestCost := math.Inf(1)
	lowestCostNode := ""

	for _, key := range keys {
		cost := costs.Get(key).Value.(float64)
		if cost < lowestCost && visitedNodes.Get(key) == nil {
			lowestCost = cost
			lowestCostNode = key
		}
	}

	return lowestCostNode
}

func getPath(predcessors *datastructures.HashTable, dest string, src string) []string {
	path := []string{dest}
	currentNode := dest

	for currentNode != "" {
		parentNode := predcessors.Get(currentNode)

		if parentNode == nil {
			break
		}

		currentNode = parentNode.Value.(string)
		path = append(path, currentNode)
	}

	return append(path, src)
}
