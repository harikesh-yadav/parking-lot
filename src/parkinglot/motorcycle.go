package parkinglot

func NewMotorcycle(licencePlate string) Vehicle {
	return &BaseVehicle{LicencePlate: licencePlate, Type: MOTORCYCLE}
}
