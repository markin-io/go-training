package algorithms

import (
	"testing"

	datastructures "github.com/markin-io/go-training/data-structures"
)

func TestBfs(t *testing.T) {
	graph := datastructures.CreateHashTable(nil)

	// Build users graph
	graph.Add("you", []string{"alice", "bob", "claire"})
	graph.Add("bob", []string{"anuj", "peggy"})
	graph.Add("alice", []string{"peggy"})
	graph.Add("claire", []string{"tom", "jonny"})
	graph.Add("anuj", []string{})
	graph.Add("peggy", []string{"you"})
	graph.Add("tom", []string{})
	graph.Add("jonny", []string{})

	searchFor := "tom"
	rootNode := "you"

	result := bfs(graph, rootNode, searchFor)

	if result != true {
		t.Errorf("bfs() not found %s", searchFor)
	}
}
