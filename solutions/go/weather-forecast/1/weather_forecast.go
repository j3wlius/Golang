// Package weather contains tools that return the weather forecast.
package weather



var (
    // CurrentCondition variable used to store the current condition.
	CurrentCondition string
    // CurrentLocation variable used to store the current location.
	CurrentLocation  string
)

// Forecast returns the current location & the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
