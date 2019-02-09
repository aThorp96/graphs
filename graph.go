package graphs

import (
    bufio
    )

type graph struct {
	numVertices int
	numEdges    int
	edges       [][]int
}

func NewGraph(numVertices int) *graph {
	g := make(graph)
	g.numVertices = numVertices
	g.clear()
	return g
}

func (g *graph) Order() int {
	return len(Edges)
}
