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

func testCountHours(t *testing.T) {
	tests := []struct {
		startTime     int
		endTime       int
		expectedHours int
	}{
		{17, 21, 4},
		{17, 24, 7},
		{18, 28, 10},
	}

	for _, testCase := range tests {
		actual := countHours(testCase.startTime, testCase.endTime)
		if actual != testCase.expectedHours {
			t.Errorf("Bad count hours from %d to %d, got=%d, want=%d", testCase.startTime, testCase.endTime, actual, testCase.expectedHours)
		}
	}
}

func TestGetHours(t *testing.T) {
	tests := []struct {
		startTime     int
		endTime       int
		expectedHours []int
	}{
		{1, 5, []int{1, 2, 3, 4}},
		{18, 20, []int{18, 19}},
		{22, 28, []int{22, 23, 24, 25, 26, 27}},
	}

	for _, testCase := range tests {
		actual := getHours(testCase.startTime, testCase.endTime)

		if !equals(actual, testCase.expectedHours) {
			t.Errorf("Bad total rate for hours from %d - %d amount... got=%d, want=%d", testCase.startTime, testCase.endTime, actual, testCase.expectedHours)
		}
	}
}

func TestFullNightCalculation(t *testing.T) {
	paymentStructure := []hourlyRate{
		{17, 23, 15},
		{23, 28, 20},
	}
	startTime := 17
	endTime := 25
	expectedTotal := (23-17)*15 + (25-23)*20
	actual := calculateTotalPayment(startTime, endTime, paymentStructure)

	if actual != expectedTotal {
		t.Errorf("Bad total payment amount... got=%d, want=%d", actual, expectedTotal)
	}
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if b[i] != e {
			return false
		}
	}

	return true
}
