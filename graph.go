package graphs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * graph is an implementation of an undirected graph.
 *
 * @author Alice McRae
 * ported to Go by Andrew Thorp
 *
 * @version 08 February 2019
 *
 *      Vertices are labeled 0..n-1, where n is
 *      the number of vertices in the graph.
 *
 *      An edge between verteces x and y is denoted
 *      as adjacencies[x][y] = true, where x > y
 */
type Undirected struct {
	// directed    bool // TODO: add directed functionality
	adjacencies [][]bool    // adjacency matrix
	edges       [][]int     // adjacency list
	weights     [][]float64 // adjacency list
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
func NewGraph(numVertices int) *Undirected {
	g := new(Undirected)
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
func NewGraphFromFile(filepath string) *Undirected {
	g := new(Undirected)
	g.Clear()
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
func NewWeightedGraphFromFile(filepath string) *Undirected {
	g := new(Undirected)
	g.Clear()
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
func (g *Undirected) readFromFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	f := bufio.NewScanner(file)
	f.Split(bufio.ScanWords)

	var vertex2 int

	f.Scan()
	g.numVertices, err = strconv.Atoi(f.Text())
	g.Clear()

	for vertex1, _ := strconv.Atoi(f.Text()); vertex1 >= 0; {

		f.Scan()
		vertex1, err = strconv.Atoi(f.Text())
		if err != nil {
			fmt.Println(err)
		}

		f.Scan()
		vertex2, err = strconv.Atoi(f.Text())
		if err != nil {
			fmt.Println(err)
		}

		if vertex1 >= 0 && vertex2 >= 0 {
			g.AddEdge(vertex1, vertex2)
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
func (g *Undirected) readWeightedFromFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	f := bufio.NewScanner(file)
	f.Split(bufio.ScanWords)

	var vertex2 int
	var weight float64

	f.Scan()
	g.numVertices, err = strconv.Atoi(f.Text())
	g.Clear()

	for vertex1, _ := strconv.Atoi(f.Text()); vertex1 >= 0; {

		f.Scan()
		vertex1, err = strconv.Atoi(f.Text())
		if err != nil {
			fmt.Println(err)
		}

		f.Scan()
		vertex2, err = strconv.Atoi(f.Text())
		if err != nil {
			fmt.Println(err)
		}

		f.Scan()
		weight, err = strconv.ParseFloat(f.Text(), 64)
		if err != nil {
			fmt.Println(err)
		}

		if vertex1 >= 0 && vertex2 >= 0 {
			g.AddEdgeWeight(vertex1, vertex2, weight)
		}
	}
}

/**
 * Accessor for the degree of a vertex.
 *
 * @param   i  vertex in the graph
 * @return  degree of vertex i
 */
func (g *Undirected) Degree(i int) int {
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
func (g *Undirected) AddEdge(vertex1, vertex2 int) {
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
func (g *Undirected) AddEdgeWeight(vertex1, vertex2 int, weight float64) {
    if vertex1 != vertex2 && !g.IsConnected(vertex1, vertex2) {
    	g.numEdges++
    	g.degrees[vertex1]++
    	g.degrees[vertex2]++

    	// inforce vertex1 > vertex2
    	if vertex1 < vertex2 {
    		temp := vertex1
    		vertex1 = vertex2
    		vertex2 = temp
    	}

    	// update
    	g.adjacencies[vertex1][vertex2] = true
    	g.weights[vertex1][vertex2] = weight
    	g.weights[vertex2][vertex1] = weight
    	g.edges[vertex1] = append(g.edges[vertex1], vertex2)
    	g.edges[vertex2] = append(g.edges[vertex2], vertex1)
    }
}

/**
 * Accessor for the connectivity of two vertices.
 *
 * @param   vertex1  vertex in the graph
 * @param   vertex2  vertex in the graph
 * @return  whether or not the vertices are connected
 */
func (g *Undirected) IsConnected(vertex1, vertex2 int) bool {
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
 *          if there is no connection 0
 */
func (g *Undirected) Weight(vertex1, vertex2 int) float64 {

	if g.adjacencies[vertex1][vertex2] || g.adjacencies[vertex2][vertex1]{
		return g.weights[vertex1][vertex2]
	} else {
		return 0
	}
}

/**
 * Accessor for the number of vertices.
 *
 * @return  number of vertices in the graph
 */
func (g *Undirected) Order() int {
	return g.numVertices
}

/**
 * Accessor for the number of edges.
 *
 * @return  number of edges in the graph
 */
func (g *Undirected) Size() int {
	return g.numEdges
}

/**
 * Accessor for edges of a vertex
 *
 * @param the vertex whos edges are to be retrieved
 *
 * @return the adjacency list of vertex
 */
func (g *Undirected) GetEdges(vertex int) []int {
	return g.edges[vertex]
}

/**
 * Removes all edges from the graph.
 */
func (g *Undirected) Clear() {
	g.numEdges = 0

	g.degrees = make([]int, g.numVertices)
	g.adjacencies = make([][]bool, g.numVertices)
	g.edges = make([][]int, g.numVertices)
	g.weights = make([][]float64, g.numVertices)

	for i := 0; i < g.numVertices; i++ {
		g.adjacencies[i] = make([]bool, g.numVertices)
		g.edges[i] = []int{}
		g.weights[i] = make([]float64, g.numVertices)
	}
}
