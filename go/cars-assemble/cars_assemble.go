package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(int(workingCarsPerHour) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const GroupsOfTenCarsCost = 95000
	const IndividualCarsCost = 10000
	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10
	return uint((groupsOfTen * GroupsOfTenCarsCost) + (remainingCars * IndividualCarsCost))
}
