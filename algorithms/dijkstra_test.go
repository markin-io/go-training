package algorithms

import (
	"testing"

	datastructures "github.com/markin-io/go-training/data-structures"
)

func TestDijkstra(t *testing.T) {
	graph := datastructures.CreateHashTable(nil)

	nodeKeys := []string{"a", "b", "c", "d", "e", "f"}

	// Create nodes
	for _, key := range nodeKeys {
		createNode(graph, key)
	}

	// Add nodes neighbors
	addNeihgbor(graph.Get("a"), "b", 5.0)
	addNeihgbor(graph.Get("a"), "c", 0.0)

	addNeihgbor(graph.Get("b"), "d", 15.0)
	addNeihgbor(graph.Get("b"), "e", 20.0)

	addNeihgbor(graph.Get("c"), "d", 30.0)
	addNeihgbor(graph.Get("c"), "e", 35.0)

	addNeihgbor(graph.Get("d"), "f", 20.0)
	addNeihgbor(graph.Get("e"), "f", 10.0)

	cost, path := dijkstra(graph, "a", "f")

	wantCost := 35.0
	wantPath := []string{"f", "e", "b", "a"}

	if wantCost != cost {
		t.Errorf("dijkstra error, incorrect cost %f, wants %f", cost, wantCost)
		return
	}

	for index, value := range path {
		if wantPath[index] != value {
			t.Errorf("dijkstra error, incorrect path %v, wants %v", path, wantPath)
			return
		}
	}
}

func createNode(graph *datastructures.HashTable, key string) {
	graph.Add(key, datastructures.CreateHashTable(nil))
}

func addNeihgbor(nodeValue *datastructures.HashNode, key string, cost float64) {
	node := nodeValue.Value.(*datastructures.HashTable)
	node.Add(key, cost)
}
