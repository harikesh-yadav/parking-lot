package parkinglot

func NewTruck(licencePlate string) Vehicle {
	return &BaseVehicle{LicencePlate: licencePlate, Type: TRUCK}
}
