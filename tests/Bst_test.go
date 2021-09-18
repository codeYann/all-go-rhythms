package tests

import (
	"testing"
	BST "github.com/codeYann/all-go-rhythms/classic"
)

func TestCreateBst(t *testing.T) {
	bst := BST.CreateTree()

	if bst.Root != nil {
		t.Error("Expected:", nil, "Got:", bst.Root)
	} else {
		t.Log("Success on create a binary search tree")
	}
}

func TestInsert(t *testing.T) {
	bst := BST.CreateTree()
	values := []int{10, 5, 12, 11, 15, 4, 9}

	for _, v := range values {
		bst.Insert(v)
	}

	result, _ := bst.Find(11)

	if result.Key != 11 {
		t.Error("Expected:", 11, "Got:", result.Key)
	} else {
		t.Log("Success on insert values")
	}
}

func TestDelte(t *testing.T) {
	bst := BST.CreateTree()
	values := []int{10, 5, 12, 11, 15, 4, 9}

	for _, v := range values {
		bst.Insert(v)
	}
}
