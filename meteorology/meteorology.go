package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tempUnit TemperatureUnit) String() string {
	var units []string = []string{"°C", "°F"}
	return units[tempUnit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp Temperature) String() string {
	return fmt.Sprint(temp.degree) + " " + temp.unit.String()
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (speedUnit SpeedUnit) String() string {
	var units []string = []string{"km/h", "mph"}
	return units[speedUnit]
}

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprint(speed.magnitude) + " " + speed.unit.String()
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteorologyData MeteorologyData) String() string {
	return meteorologyData.location + ": " +
		meteorologyData.temperature.String() +
		", Wind " + meteorologyData.windDirection +
		" at " + meteorologyData.windSpeed.String() +
		", " + fmt.Sprint(meteorologyData.humidity) + "% Humidity"
}
