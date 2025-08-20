//Package weather Provide method To show current forecast of a city.
package weather
// CurrentCondition location of the forcast. 
var CurrentCondition string
// CurrentLocation, location of the forecast.
var CurrentLocation string
// Forecast return the forecast in a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
