// Package weather tells current weather condition on the location.
package weather

// CurrentCondition current weather.
var CurrentCondition string
// CurrentLocation keeps current location.
var CurrentLocation string

// Forecast return: string with message about current weather on the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
