package tests

import (
	"testing"

	BST "github.com/codeYann/all-go-rhythms/classic"
)

func TestCreateBst(t *testing.T) {
	bst := BST.CreateTree()

	if bst.GetRoot() != nil {
		t.Error("Expected:", nil, "Got:", bst.GetRoot())
	} else {
		t.Log("Success on create a binary search tree")
	}
}

func TestInsertBst(t *testing.T) {
	bst := BST.CreateTree()

	bst.Insert(10)
	bst.Insert(15)
	bst.Insert(8)

	if bst.GetRoot().Key != 10 {
		t.Error("Expected:", 10, "Got:", bst.GetRoot().Key)
	}

	if bst.GetLength() != 3 {
		t.Error("Expected:", 3, "Got:", bst.GetLength())
	}
}

func TestHeight(t *testing.T) {
	bst := BST.CreateTree()

	bst.Insert(30)
	bst.Insert(25)
	bst.Insert(40)
	bst.Insert(28)
	bst.Insert(15)
	bst.Insert(6)
	bst.Insert(9)
	bst.Insert(39)
	bst.Insert(42)

	height := bst.Height()

	if height != 4 {
		t.Error("Expected:", 4, "Got:", height)
	}
}
