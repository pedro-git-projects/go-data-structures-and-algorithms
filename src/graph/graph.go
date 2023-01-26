package graph

import (
	"fmt"

	"github.com/pedro-git-projects/go-data-structures-and-algorithms/src/errs"
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

func (g *Graph[T]) AddVertex(vertex T) error {
	_, present := g.adjacencyList[vertex]
	if !present {
		g.adjacencyList[vertex] = set.New[T]()
		return nil
	}
	return errs.DuplicatedVertex(vertex)
}

func (g *Graph[T]) AddEdge(v0, v1 T) error {
	_, present := g.adjacencyList[v0]
	if !present {
		return errs.NonExistentVertex(v0)
	}

	_, present = g.adjacencyList[v1]
	if !present {
		return errs.NonExistentVertex(v1)
	}

	g.adjacencyList[v0].Insert(v1)
	g.adjacencyList[v1].Insert(v0)
	return nil
}

func (g *Graph[T]) RemoveEdge(v0, v1 T) error {
	_, present := g.adjacencyList[v0]
	if !present {
		return errs.NonExistentVertex(v0)
	}
	_, present = g.adjacencyList[v1]
	if !present {
		return errs.NonExistentVertex(v1)
	}

	g.adjacencyList[v0].Erase(v1)
	g.adjacencyList[v1].Erase(v0)
	return nil
}

func (g *Graph[T]) RemoveVertex(v T) error {
	_, present := g.adjacencyList[v]
	if !present {
		return errs.NonExistentVertex(v)
	}

	set := g.adjacencyList[v]
	setSlice := set.ToSlice()

	for _, otherVertexes := range setSlice {
		g.adjacencyList[otherVertexes].Erase(v)
	}
	delete(g.adjacencyList, v)

	return nil
}
