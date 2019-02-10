package graphs

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	//"io"
)

type graph struct {
	directed    bool
	numVertices int
	numEdges    int
	adjacencies [][]bool // adjacency matrix
}

func NewGraph(numVertices int) *graph {
	g := new(graph)
	g.numVertices = numVertices
	g.Clear()
	return g
}

func NewGraphFromFile(filepath string) *graph {
	g := new(graph)
	g.readFromFile(filepath)
	return g
}

func (g *graph) readFromFile(filepath string) {

	f := bufio.NewScanner(strings.NewReader(filepath))
	f.Split(bufio.ScanWords)

	var vertex1, vertex2 int
	var err error

	f.Scan()
	g.numVertices, err = strconv.Atoi(f.Text())
	g.Clear()

	for b, _ := strconv.Atoi(f.Text()); b >= 0; {
		f.Scan()
		vertex1, err = strconv.Atoi(f.Text())
		f.Scan()
		vertex2, err = strconv.Atoi(f.Text())
		g.AddEdge(vertex1, vertex2)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (g *graph) AddEdge(vertex1, vertex2 int) {
	g.numEdges++
	if vertex1 > vertex2 {
		g.adjacencies[vertex1][vertex2] = true
	} else {
		g.adjacencies[vertex2][vertex1] = true
	}
}

func (g *graph) isEdge(vertex1, vertex2 int) bool {
	if vertex1 > vertex2 {
		return g.adjacencies[vertex1][vertex2]
	} else {
		return g.adjacencies[vertex2][vertex1]
	}
}

func (g *graph) Order() int {
	return g.numVertices
}

func (g *graph) Size() int {
	return g.numEdges
}

func (g *graph) Clear() {
	g.numEdges = 0
	g.adjacencies = make([][]bool, g.numVertices)
	for i := 0; i < g.numVertices; i++ {
		g.adjacencies[i] = make([]bool, g.numVertices)
	}
}
