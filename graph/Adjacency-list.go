package graph

type Graph struct {
	Vertices map[int][]int
}

func CreateGraph() *Graph {
	return &Graph{Vertices: make(map[int][]int)}
}

func (g *Graph) GetEdge(source int) []int {
	return g.Vertices[source]
}

func (g *Graph) Append(a, b int) {
	g.Vertices[a] = append(g.Vertices[a], b)
	g.Vertices[b] = append(g.Vertices[b], a)
}

func (d *Graph) AppendDigraph(a, b int) {
	d.Vertices[a] = append(d.Vertices[a], b)
}

func (g *Graph) GetVerticesNumbers() int {
	return len(g.Vertices)
}
