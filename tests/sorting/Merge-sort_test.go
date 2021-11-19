package tests

import (
	"testing"

	s "github.com/codeYann/all-go-rhythms/sorting"
)

func TestMergeSort(t *testing.T) {
	A := []int{90, 45, 10, 15, 21, 5, 4, 3, 1}
	B := s.MergeSort(A)

	if B[0] != 1 {
		t.Error("Expected: ", 1, "Got:", B[0])
	}
}
