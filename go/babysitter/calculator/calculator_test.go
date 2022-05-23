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

func TestGetRate(t *testing.T) {
	paymentStructure := []hourlyRate{
		{17, 23, 15},
		{23, 26, 20},
		{26, 28, 50},
	}

	tests := []struct {
		hour         int
		expectedRate int
	}{
		{17, 15},
		{19, 15},
		{23, 20},
		{25, 20},
		{26, 50},
		{27, 50},
		{28, 0}, //wages are not defined for the 4am-5am time window
	}

	for _, testCase := range tests {
		actual := getRate(paymentStructure, testCase.hour)

		if actual != testCase.expectedRate {
			t.Errorf("Bad total rate for hour %d amount... got=%d, want=%d", testCase.hour, actual, testCase.expectedRate)
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
