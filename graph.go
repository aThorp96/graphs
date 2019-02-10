package graphs

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	//"io"
)

/**
 * graph is an implementation of an undirected graph.
 *
 * @author Alice McRae
 * ported to Go by
 * @author Andrew Thorp
 * @version 08 February 2019
 *
 *      Vertices are labeled 0..n-1, where n is
 *      the number of vertices in the graph.
 *
 *      An edge between verteces x and y is denoted
 *      as adjacencies[x][y] = true, where x > y
 */
type graph struct {
	// directed    bool // TODO: add directed functionality
	adjacencies [][]bool // adjacency matrix
	degrees     []int
	numVertices int
	numEdges    int
}

/**
 * Constructor sets up the adjacency lists for a graph
 *       with a set number of vertices, and no edges
 *
 * @param num  number of vertices in the graph
 */
func NewGraph(numVertices int) *graph {
	g := new(graph)
	g.numVertices = numVertices
	g.Clear()
	return g
}

/**
 * Constructor sets up the adjacency lists for a graph
 *       from a file.  The file is in the format
 *       first entry: the number of vertices
 *       subsequent entries: pairs of vertices
 *               representing the edges
 *
 * @param filename  name of the input file
 */
func NewGraphFromFile(filepath string) *graph {
	g := new(graph)
	g.readFromFile(filepath)
	return g
}

/**
 * Inputs adjacency lists from a file.
 *
 * @param filename  name of the input file
 *
 *       Reads the number of vertices and
 *       each edge from a file.  The file format is
 *       first entry: the number of vertices
 *       subsequent entries: pairs of vertices
 *                          representing the edges.
 */
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

/**
* Accessor for the degree of a vertex.
*
* @param   i  vertex in the graph
* @return  degree of vertex i
 */
func (g *graph) Degree(i int) int {
	return g.degrees[i]
}

/**
 * Adds an edge uv to an undirected graph.
 *
 * @param vertex1  one endpoint
 * @param vertex2  one endpoint
 *
 * The smaller of the inputs is added to the larger
 * vertice's list
 */
func (g *graph) AddEdge(vertex1, vertex2 int) {
	g.numEdges++
	g.degrees[vertex1]++
	g.degrees[vertex2]++

	if vertex1 > vertex2 {
		g.adjacencies[vertex1][vertex2] = true
	} else {
		g.adjacencies[vertex2][vertex1] = true
	}
}

/**
* Accessor for the connectivity of two vertices.
*
* @param   vertex1  vertex in the graph
* @param   vertex2  vertex in the graph
* @return  whether or not the vertices are connected
 */
func (g *graph) IsEdge(vertex1, vertex2 int) bool {
	if vertex1 > vertex2 {
		return g.adjacencies[vertex1][vertex2]
	} else {
		return g.adjacencies[vertex2][vertex1]

	}
}

/**
 * Accessor for the number of vertices.
 *
 * @return  number of vertices in the graph
 */
func (g *graph) Order() int {
	return g.numVertices
}

/**
 * Accessor for the number of edges.
 *
 * @return  number of edges in the graph
 */
func (g *graph) Size() int {
	return g.numEdges
}

/**
 * Removes all edges from the graph.
 */
func (g *graph) Clear() {
	g.numEdges = 0
	g.degrees = make([]int, g.numVertices)
	g.adjacencies = make([][]bool, g.numVertices)
	for i := 0; i < g.numVertices; i++ {
		g.adjacencies[i] = make([]bool, g.numVertices)
	}
}
