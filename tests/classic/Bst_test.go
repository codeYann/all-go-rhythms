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
