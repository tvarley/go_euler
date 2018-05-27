package go_euler

import "testing"

func TestSolution004(t *testing.T) {
	expected := 906609
	actual := Solution004()
	if actual != expected {
		t.Errorf("Solution004 == %d, expected %d", actual, expected)
	}
}
