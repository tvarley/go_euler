package go_euler

import "testing"

func TestSolution002(t *testing.T) {
	expected := 4613732
	actual := Solution002()
	if actual != expected {
		t.Errorf("Solution002 == %d, expected %d", actual, expected)
	}
}
