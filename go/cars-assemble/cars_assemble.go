package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (0.01 * successRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carFloat := (float64(productionRate) * (0.01 * successRate)) / 60
	return int(carFloat)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// nCars := 21
	// discount := (nCars / 10) * 95000
	// orig := (nCars % 10) * 10000
	// price := discount + orig
	return uint((carsCount/10)*95000 + (carsCount%10)*10000)
}
