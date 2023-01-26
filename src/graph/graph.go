package graph

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/set"
)

type Graph[T comparable] struct {
	adjacencyList map[T]set.Set[T]
}

func New[T comparable](vertexID T) *Graph[T] {
	m := make(map[T]set.Set[T])
	m[vertexID] = set.New[T]()
	vertex := m
	g := &Graph[T]{
		adjacencyList: vertex,
	}
	return g
}

func (g Graph[T]) AdjecencyList() map[T]set.Set[T] {
	return g.adjacencyList
}

func (g Graph[T]) String() string {
	var str string
	for k, v := range g.adjacencyList {
		str += fmt.Sprintf("%v: [ %v ]\n", k, v)
	}
	return str
}

func (g *Graph[T]) AddVertex(vertex T) bool {
	_, ok := g.adjacencyList[vertex]
	if !ok {
		g.adjacencyList[vertex] = set.New[T]()
		return true
	}
	return false
}

func (g *Graph[T]) AddEdge(v0, v1 T) bool {
	_, ok0 := g.adjacencyList[v0]
	_, ok1 := g.adjacencyList[v1]

	if ok0 && ok1 {
		g.adjacencyList[v0].Insert(v1)
		g.adjacencyList[v1].Insert(v0)
		return true
	}

	return false
}
