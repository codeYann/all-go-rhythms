package tests

import (
	"testing"

	C "github.com/codeYann/all-go-rhythms/classic"
)

func TestCreateCircularLinkedList(t *testing.T) {
	ciruclar := C.CreateCircularLinkedList()
	if ciruclar.GetSize() != 0 {
		t.Error("Expected:", 0, "Got:", ciruclar.GetSize())
	}
}

func TestAppendCircularLinkedList(t *testing.T) {
	ciruclar := C.CreateCircularLinkedList()
	ciruclar.Append(10)
	ciruclar.Append(15)
	ciruclar.Append(30)

	if ciruclar.GetTail().Key != 30 {
		t.Error("Expected:", 30, "Got:", ciruclar.GetTail().Key)
	}

	if ciruclar.GetTailNext().Key != 10 {
		t.Error("Expected:", 10, "Got:", ciruclar.GetTailNext().Key)
	}
}
