package numbers

import "testing"

func TestSimilarity(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	similarity := CalcSimilarity(a, b)
	expected := 31
	if similarity != expected {
		t.Fatalf("Expected similarity of %v but got %v", expected, similarity)
	}
}
