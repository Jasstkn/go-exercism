package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / 100 * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(successRate / 100 * float64(productionRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	const tenCost = 95000
	const normalCost = 10000
	return uint(carsCount) / 10 * tenCost + uint(carsCount) % 10 * normalCost
}
