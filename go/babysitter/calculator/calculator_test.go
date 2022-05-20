package calculator

import "testing"

func TestCalculate(t *testing.T) {
	expected := 1
	actual := calculate()
	if actual != expected {
		t.Errorf("actual %d does not match expected %d", actual, expected)
	}
}
