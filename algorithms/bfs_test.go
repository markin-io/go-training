package algorithms

import (
	"testing"

	datastructures "github.com/markin-io/go-training/data-structures"
)

func TestBfs(t *testing.T) {
	graph := datastructures.CreateHashTable(nil)

	// Build users graph
	graph.Add("0", []string{"1", "3"})
	graph.Add("1", []string{"2", "0"})
	graph.Add("2", []string{"1"})
	graph.Add("3", []string{"4", "7"})
	graph.Add("4", []string{"5", "6", "7"})
	graph.Add("5", []string{"6"})
	graph.Add("6", []string{"7", "8"})
	graph.Add("7", []string{})
	graph.Add("8", []string{})

	shortestPath := []string{"7", "3", "0"}
	distance := 2

	src := "0"
	dest := "7"

	if bfsIsPathExists(graph, src, dest) == false {
		t.Errorf("bfs error, path %s -> %s exists", src, dest)
		return
	}

	foundShortestPath, foundDistance := bfsFindPath(graph, src, dest)

	if foundDistance != distance {
		t.Errorf("bfs incorrect shortest path %s", dest)
	}

	for index, value := range foundShortestPath {
		if shortestPath[index] != value {
			t.Errorf("bfs incorrect shortest path. Expects %v, got %v", shortestPath, foundShortestPath)
			break
		}
	}
}
