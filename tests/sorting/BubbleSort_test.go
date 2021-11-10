package tests

import (
	"testing"

	s "github.com/codeYann/all-go-rhythms/sorting"
)

func TestBubbleSort(t *testing.T) {
	Vetor := []int{10, 15, 30, 45, 80, 1}
	s.BubbleSort(Vetor)

	if Vetor[0] != 1 {
		t.Error("Expected:", 1, "Got:", Vetor[0])
	}
}
