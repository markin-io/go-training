package algorithms

import (
	"testing"

	datastructures "github.com/markin-io/go-training/data-structures"
)

func TestDijkstra(t *testing.T) {
	// graph := datastructures.CreateHashTable(nil)

	// // Initialize nodes
	// graph.Add("start", datastructures.CreateHashTable(nil))
	// graph.Add("a", datastructures.CreateHashTable(nil))
	// graph.Add("b", datastructures.CreateHashTable(nil))
	// graph.Add("fin", datastructures.CreateHashTable(nil))

	// // Add nodes neighbors
	// graph.Get("start").Value.(*datastructures.HashTable).Add("a", 6.0)
	// graph.Get("start").Value.(*datastructures.HashTable).Add("b", 2.0)

	// graph.Get("a").Value.(*datastructures.HashTable).Add("fin", 1.0)
	// graph.Get("b").Value.(*datastructures.HashTable).Add("a", 3.0)
	// graph.Get("b").Value.(*datastructures.HashTable).Add("fin", 5.0)

	// // Initialize costs
	// costs := datastructures.CreateHashTable(nil)
	// costs.Add("a", 6.0)
	// costs.Add("b", 2.0)
	// costs.Add("fin", math.Inf(1))

	// cost, path := dijkstra(graph, costs, "start", "fin")

	// log.Printf("Cost %v %v", cost, path)

	graph := datastructures.CreateHashTable(nil)

	nodeKeys := []string{"a", "b", "c", "d", "e", "f"}

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

	wantCost := 35.0
	wantPath := []string{"f", "e", "b", "a"}

	cost, path := dijkstra(graph, "a", "f")

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
