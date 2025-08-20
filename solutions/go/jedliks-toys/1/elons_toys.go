package elon

import (
	"fmt"
)

func (car *Car) Drive() {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%s", car.battery, "%")
}

func (car *Car) CanFinish(trackDistance int) bool {
	remainingDistance := trackDistance - car.distance
	estimatedDrain := (remainingDistance / car.speed) * car.batteryDrain
	return  estimatedDrain <= car.battery
}
