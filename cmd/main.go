package main

import "github.com/parking-lot/src/parkinglot"

func main() {

	parking := parkinglot.GetParkingLotInstane()
	level := parkinglot.NewLevel(1, 5)
	parking.AddLevel(level)

	car := parkinglot.NewCar("24BHCAR")
	bike := parkinglot.NewMoterCycle("24BH8149N")
	bike2 := parkinglot.NewMoterCycle("25BHBIKE")
	parking.ParkVehicle(car)
	parking.ParkVehicle(bike)
	parking.ParkVehicle(bike2)
	parking.DisplayAvailability()

}
