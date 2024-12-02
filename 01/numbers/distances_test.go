package numbers

import "testing"

func TestCalculateDistance(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	distance, err := CalcDistances(a, b)
	if err != nil {
		t.Fatalf("Unexpected error calculating distances: %v", err)
	}

	expected_distance := 11
	if distance != expected_distance {
		t.Fatalf("Expected distance of %v but got %v", expected_distance, distance)
	}
}
