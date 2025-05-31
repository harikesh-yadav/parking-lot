package main

import "github.com/parking-lot/src/parkinglot"

func main() {

	parking := parkinglot.GetParkingLotInstance()
	parking.AddLevel(parkinglot.NewLevel(1, 4))
	parking.AddLevel(parkinglot.NewLevel(2, 2))

	car := parkinglot.NewCar("ABC123")
	truck := parkinglot.NewTruck("XYZ789")
	motorcycle := parkinglot.NewMotorcycle("M1234")

	parking.ParkVehicle(car)
	parking.ParkVehicle(truck)
	parking.ParkVehicle(motorcycle)

	parking.UnParkVehicle(motorcycle)

	parking.DisplayAvailability()

}
