package go_euler

import "testing"

func TestSolution001(t *testing.T) {
	expected := 233168
	actual := Solution001()
	if actual != expected {
		t.Errorf("Solution001 == %d, expected %d", actual, expected)
	}
}
