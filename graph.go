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
	edges       [][]int  // adjacency list
	weights     [][]int  // adjacency list
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
 * Constructor sets up the adjacency lists for a graph
 *       from a file.  The file is in the format
 *       first entry: the number of vertices
 *       subsequent entries: pairs of vertices
 *               representing the edges followed
 *               by the weight of the vertex pair edge
 *
 * @param filename  name of the input file
 */
func NewWeightedGraphFromFile(filepath string) *graph {
	g := new(graph)
	g.readWeightedFromFile(filepath)
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
func (g *graph) readWeightedFromFile(filepath string) {

	f := bufio.NewScanner(strings.NewReader(filepath))
	f.Split(bufio.ScanWords)

	var vertex1, vertex2, weight int
	var err error

	f.Scan()
	g.numVertices, err = strconv.Atoi(f.Text())
	g.Clear()

	for b, _ := strconv.Atoi(f.Text()); b >= 0; {
		f.Scan()
		vertex1, err = strconv.Atoi(f.Text())
		f.Scan()
		vertex2, err = strconv.Atoi(f.Text())
		f.Scan()
		weight, err = strconv.Atoi(f.Text())
		g.AddEdgeWeight(vertex1, vertex2, weight)

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
	g.AddEdgeWeight(vertex1, vertex2, 1)
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
func (g *graph) AddEdgeWeight(vertex1, vertex2 int, weight int) {
	g.numEdges++
	g.degrees[vertex1]++
	g.degrees[vertex2]++

	if vertex1 < vertex2 {
		temp := vertex1
		vertex1 = vertex2
		vertex2 = temp
	}
	g.adjacencies[vertex1][vertex2] = true
	g.weights[vertex1][vertex2] = weight
	g.edges[vertex1] = append(g.edges[vertex1], vertex2)

}

/**
 * Accessor for the connectivity of two vertices.
 *
 * @param   vertex1  vertex in the graph
 * @param   vertex2  vertex in the graph
 * @return  whether or not the vertices are connected
 */
func (g *graph) IsConnected(vertex1, vertex2 int) bool {
	if vertex1 > vertex2 {
		return g.adjacencies[vertex1][vertex2]
	} else {
		return g.adjacencies[vertex2][vertex1]

	}
}

/**
 * Accessor for the weight of an edge.
 *
 * @param   vertex1  vertex in the graph
 * @param   vertex2  vertex in the graph
 * @return  the weight of the connected edge
 *          if there is no connection -999
 */
func (g *graph) Weight(vertex1, vertex2 int) int {
	if vertex2 > vertex1 {
		temp := vertex1
		vertex1 = vertex2
		vertex2 = temp
	}
	if g.adjacencies[vertex1][vertex2] {
		return g.weights[vertex1][vertex2]
	} else {
		return -999
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
 * Accessor for edges of a vertex
 *
 * @param the vertex whos edges are to be retrieved
 *
 * @return the adjacency list of vertex
 */
func (g *graph) GetEdges(vertex int) []int {
	return g.edges[vertex]
}

/**
 * Removes all edges from the graph.
 */
func (g *graph) Clear() {
	g.numEdges = 0

	g.degrees = make([]int, g.numVertices)
	g.adjacencies = make([][]bool, g.numVertices)
	g.edges = make([][]int, g.numVertices)
	g.weights = make([][]int, g.numVertices)

	for i := 0; i < g.numVertices; i++ {
		g.adjacencies[i] = make([]bool, g.numVertices)
		g.edges[i] = make([]int, g.numVertices)
		g.weights[i] = make([]int, g.numVertices)
	}
}
