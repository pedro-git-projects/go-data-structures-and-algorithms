package graph_test

import (
	"fmt"
	"testing"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/graph"
)

func TestAddVertex(t *testing.T) {
	g := graph.New("A", 1)
	g.AddVertex("B")
	fmt.Println(g)
}
