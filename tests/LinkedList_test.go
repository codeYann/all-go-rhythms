package tests

import (
	"testing"

	List "github.com/codeYann/all-go-rhythms/classic"
)

func TestCreateLinkedList(t *testing.T) {
	list := List.CreateLinkedList()
	if list.Size() != 0 {
		t.Error("Expected:", 0, "Got:", list.Size())
	} else {
		t.Log("Success on create Linked List")
	}
}

func TestAddBack(t *testing.T) {
	list := List.CreateLinkedList()
	values := []int{1, 2, 4, 8, 16, 32}

	for _, v := range values {
		list.AddBack(v)
	}

	if list.GetTail().Key != 32 {
		t.Error("Expected:", 32, "Got:", list.GetTail().Key)
	} else {
		t.Log("Success on add back")
	}
}

func TestAddFront(t *testing.T) {
	list := List.CreateLinkedList()
	values := []int{1, 2, 4, 8, 16, 32}

	for _, v := range values {
		list.AddBack(v)
	}

	if list.GetHead().Key != 1 {
		t.Error("Expected:", 1, "Got:", list.GetHead().Key)
	} else {
		t.Log("Success on add front")
	}
}

func TestRemoveBack(t *testing.T) {
	list := List.CreateLinkedList()
	values := []int{1, 2, 4, 8, 16, 32}

	for _, v := range values {
		list.AddBack(v)
	}

	node, _ := list.RemoveBack()

	if node.Key != 32 {
		t.Error("Expected:", 32, "Got:", node.Key)
	}

	if list.Size() != 5 {
		t.Error("Expected:", 5, "Got:", list.Size())
	}

	if list.GetTail().Key != 16 {
		t.Error("Expected:", 16, "Got:", list.GetTail().Key)
	}
}

func TestRemoveFront(t *testing.T) {
	list := List.CreateLinkedList()
	values := []int{1, 2, 4, 8, 16, 32}

	for _, v := range values {
		list.AddBack(v)
	}

	node, _ := list.RemoveFront()

	if node.Key != 1 {
		t.Error("Expected:", 1, "Got:", node.Key)
	}

	if list.Size() != 5 {
		t.Error("Expected:", 5, "Got:", list.Size())
	}

	if list.GetHead().Key != 2 {
		t.Error("Expected:", 2, "Got:", list.GetHead().Key)
	}
}
