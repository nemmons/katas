package calculator

type hours []int

func calculatePayment(hourlyRate int, hours int) int {
	return hourlyRate * hours
}

// Here startTime and endTime represent time as integer hours in 24-hour time,
//with values >  24 representing the next morning (e.g. 17 = 5pm, 26 = 2am)
func countHours(startTime int, endTime int) int {
	return endTime - startTime
}

type hourlyRate struct {
	startTime int
	endTime   int
	rate      int
}

func calculateTotalPayment(startTime int, endTime int, paymentStructure []hourlyRate) int {
	total := 0

	workedHours := getHours(startTime, endTime)

	for _, entry := range paymentStructure {
		count := workedHours.countFrom(entry.startTime, entry.endTime)
		total += entry.rate * count
	}

	return total
}

func getHours(startTime int, endTime int) hours {
	var result hours
	for i := startTime; i < endTime; i++ {
		result = append(result, i)
	}
	return result
}

func filter(data []int, f func(int) bool) []int {
	var output []int

	for _, x := range data {
		if f(x) == true {
			output = append(output, x)
		}
	}

	return output
}

func (h hours) countFrom(startTime int, endTime int) int {
	return len(filter(h, func(hour int) bool {
		return startTime <= hour && hour < endTime
	}))
}
