package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	return maxDistance(car) >= trackDistance
}

func maxDistance(car Car) int {
	if car.battery < car.batteryDrain {
		return 0
	}
	car.battery -= car.batteryDrain
	return car.speed + maxDistance(car)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
