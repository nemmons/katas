package calculator

import "testing"

func TestCalculateFromHourlyRate(t *testing.T) {
	tests := []struct {
		hourlyRate    int
		hours         int
		expectedTotal int
	}{
		{15, 5, 75},
		{12, 1, 12},
		{100, 0, 0},
	}

	for _, testCase := range tests {
		actual := calculatePayment(testCase.hourlyRate, testCase.hours)
		if actual != testCase.expectedTotal {
			t.Errorf("Bad calculation for %d hours at an hourly rate of %d, got=%d, want=%d", testCase.hours, testCase.hourlyRate, actual, testCase.expectedTotal)
		}
	}
}
