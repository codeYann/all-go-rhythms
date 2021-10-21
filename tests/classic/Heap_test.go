package tests

import (
	"testing"

	H "github.com/codeYann/all-go-rhythms/classic"
)

func TestCreateMaxHeap(t *testing.T) {
	heap := H.CreateMaxHeap()

	if heap.Size() != 0 {
		t.Error("Expected:", 0, "Got:", heap.Size())
	}

}

func TestInsertMaxHeap(t *testing.T) {
	heap := H.CreateMaxHeap()

	heap.Insert(90)
	heap.Insert(100)
	heap.Insert(80)
	heap.Insert(72)
	heap.Insert(177)
	heap.Insert(73)

	if heap.GetElements()[0] != 177 {
		t.Error("Expected:", 177, "Got:", heap.GetElements()[0])
	}
}
