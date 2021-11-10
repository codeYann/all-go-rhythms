package tests

import (
	"testing"

	s "github.com/codeYann/all-go-rhythms/sorting"
)

func TestCountingSort(t *testing.T) {
	Vec := []int{5, 5, 2, 8, 10, 1}
	a := s.CountSort(Vec)

	if a[0] != 1 {
		t.Error("Expected:", 1, "Got:", a[0])
	}
}
