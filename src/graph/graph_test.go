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
	err := g.AddEdge(1, 2)

	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestRemoveEdge(t *testing.T) {
	g := graph.New(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 3)

	err := g.RemoveEdge(1, 2)
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	err = g.RemoveEdge(4, 2)
	if err == nil {
		t.Error("Expected an error but got none")
	}
}
