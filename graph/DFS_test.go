package graph

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	g := CreateGraph()
	g.Append(0, 1)
	g.Append(0, 2)
	g.Append(2, 3)
	g.Append(3, 1)

	infos := DFS(g)
	fmt.Println(infos)

	if infos.D[0] != 1 {
		t.Error("Expected:", 1, "Got:", infos.D[0])
	} else {
		t.Log("Success")
	}
}
