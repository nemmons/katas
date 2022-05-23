package calculator

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
	//TODO this lacks elegance.
	for i := startTime; i < endTime; i++ {
		total += getRate(paymentStructure, i)
	}
	return total
}

func getRate(paymentStructure []hourlyRate, hour int) int {
	for _, r := range paymentStructure {
		if r.startTime <= hour && hour < r.endTime {
			return r.rate
		}
	}
	return 0 //TODO update this method to return a status, result tuple and return error here
}
