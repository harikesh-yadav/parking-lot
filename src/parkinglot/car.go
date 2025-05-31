package parkinglot

func NewCar(licencePlate string) Vehicle {
	return &BaseVehicle{LicencePlate: licencePlate, Type: CAR}
}
