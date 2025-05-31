package parkinglot

func NewMoterCycle(licencePlate string) Vehicle {
	return &BaseVehicle{LicencePlate: licencePlate, Type: MOTORCYCLE}
}
