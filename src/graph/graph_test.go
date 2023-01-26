package graph_test

import (
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/graph"
)

func TestAddVertex(t *testing.T) {
	g := graph.New("A")
	got := len(g.AdjecencyList())
	expect := 1

	if got != expect {
		t.Errorf("Expected %v but got %v", expect, got)
	}

	g.AddVertex("B")
	got = len(g.AdjecencyList())
	expect = 2

	if got != expect {
		t.Errorf("Expected %v but got %v", expect, got)
	}
}

func TestAddEdge(t *testing.T) {
	g := graph.New(1)
	g.AddVertex(2)
	got := g.AddEdge(1, 2)
	expect := true

	if got != expect {
		t.Errorf("Expected %v but got %v", expect, got)
	}
}
