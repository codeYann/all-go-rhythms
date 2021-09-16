package graph

import "fmt"

type graphInfo struct {
	colors []string
	D      []int
	F      []int
}

var mark int = 0

func DFS_VISIT(lista_adj map[int][]int, vertice int, markPointer *int, infos *graphInfo) {
	fmt.Println(vertice)
	*markPointer += 1

	infos.D[vertice] = *markPointer
	infos.colors[vertice] = "C"

	for _, u := range lista_adj[vertice] {
		if infos.colors[vertice] == "B" {
			DFS_VISIT(lista_adj, u, markPointer, infos)
		}
	}

	*markPointer += 1
	infos.F[vertice] = *markPointer
	infos.colors[vertice] = "P"
}

func DFS(g *Graph) *graphInfo {
	markValue := &mark
	*markValue = 0

	infos := &graphInfo{
		colors: make([]string, 0),
		D:      make([]int, 0),
		F:      make([]int, 0),
	}

	for i := 0; i < g.GetVerticesNumbers(); i++ {
		infos.colors = append(infos.colors, "B")
		infos.D = append(infos.D, 0)
		infos.F = append(infos.F, 0)
	}

	for v := range g.GetAllVertices() {
		if infos.colors[v] == "B" {
			DFS_VISIT(g.Vertices, v, &mark, infos)
		}
	}

	return infos
}
