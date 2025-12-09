package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
    return float64(productionRate) * (successRate * 0.01)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
    if(successRate == 0 || successRate == 0.0){ 
        return 0
    } else {
        return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
    }
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
    var countDividedByTen = carsCount / 10
    var reminderWhenDividedByTen = carsCount % 10

    return uint((countDividedByTen * 95000) + (reminderWhenDividedByTen * 10000))
}
