package tests

import (
	"testing"

	s "github.com/codeYann/all-go-rhythms/sorting"
)

func TestQuickSort(t *testing.T) {
	Vetor := []int{10, 15, 30, 45, 80, 1}
	s.QuickSort(Vetor, 0, 5)
	if Vetor[0] != 1 {
		t.Error("Expected:", 1, "Got:", Vetor[0])
	}

	if Vetor[5] != 80 {
		t.Error("Expected:", 80, "Got:", Vetor[5])
	}
}
