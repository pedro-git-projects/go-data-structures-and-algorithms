package graph

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/set"
)

type Graph[T, U comparable] struct {
	adjacencyList map[T]set.Set[U]
}

func New[T, U comparable](vertexID T, vertexValue U) *Graph[T, U] {
	m := make(map[T]set.Set[U])
	s := set.New(vertexValue)
	m[vertexID] = s
	vertex := m
	g := &Graph[T, U]{
		adjacencyList: vertex,
	}
	return g
}

func (g Graph[T, U]) String() string {
	var str string
	for k, v := range g.adjacencyList {
		str += fmt.Sprintf("%v: [ %v ]\n", k, v)
	}
	return str
}

func (g *Graph[T, U]) AddVertex(vertex T) bool {
	_, ok := g.adjacencyList[vertex]
	if !ok {
		g.adjacencyList[vertex] = set.New[U]()
		return true
	}
	return false
}

// func (g *Graph[T, U]) AddEdge(vertex0, vertex1 U) {
// }
