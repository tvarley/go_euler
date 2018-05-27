package go_euler

import "testing"

func TestSolution003(t *testing.T) {
	expected := 6857
	actual := Solution003()
	if actual != expected {
		t.Errorf("Solution003 == %d, expected %d", actual, expected)
	}
}
