package main

import (
	"fmt"
	"log"
)

//graph
type Graph struct {
	vertices []*Vertex
}

//vertex
type Vertex struct {
	key int
	adj []*Vertex //TODO: make linked list
}

//add vertex
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		log.Printf("key %v already exists\n", k)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//add edge (directed)
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		log.Printf("invalid edge %v->%v\n", from, to)
		return
	} else if contains(fromVertex.adj, toVertex.key) {
		log.Println("edge already exists: from", from, "to", to)
		return
	}

	fromVertex.adj = append(fromVertex.adj, toVertex)
}

//pretty print
func (g *Graph) PrettyPrint() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %v: ", v.key)
		for _, v := range v.adj {
			fmt.Printf("%v ", v.key)
		}
		fmt.Println("")
	}
}

//get vertex
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

//contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddVertex(0)
	test.AddEdge(1, 2)
	test.AddEdge(3, 5)
	test.AddEdge(1, 2)

	test.PrettyPrint()
}
