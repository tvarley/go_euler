package go_euler

import "testing"

func TestSolution005(t *testing.T) {
	expected := 232792560
	actual := Solution005()
	if actual != expected {
		t.Errorf("Solution005 == %d, expected %d", actual, expected)
	}
}
