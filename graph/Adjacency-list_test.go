package graph

import (
	"testing"
)

func TestAppendGraph(t *testing.T) {
	G := CreateGraph()
	G.Append(0, 1)
	G.Append(1, 2)
	G.Append(2, 0)

	a := G.GetEdge(0)
	b := G.GetEdge(1)

	if a[0] != 1 {
		t.Error("Expected:", 1, "Got:", a[0])
	} else {
		t.Log("Success")
	}

	if b[1] != 2 {
		t.Error("Expected:", 2, "Got:", b[1])
	} else {
		t.Log("Success")
	}
}

func TestAppendDigraph(t *testing.T) {
	d := CreateGraph()

	numbers := []int{2, 4, 8, 16, 32, 64}

	for i := 0; i < len(numbers)-1; i++ {
		d.AppendDigraph(numbers[i], numbers[i+1])
	}

	v := d.GetEdge(4)

	if v[0] != numbers[2] {
		t.Error("Expected:", 8, "Got:", v[2])
	} else {
		t.Log("Success on add edges in Digraph")
	}
}
