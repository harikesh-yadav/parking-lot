package parkinglot

type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

type Vehicle interface {
	GetLicencePlate() string
	GetType() VehicleType
	VehicleName() string
}

type BaseVehicle struct {
	LicencePlate string
	Type         VehicleType
}

func (b *BaseVehicle) GetLicencePlate() string {
	return b.LicencePlate
}
func (b *BaseVehicle) GetType() VehicleType {
	return b.Type
}
func (b *BaseVehicle) VehicleName() string {
	if b.Type == CAR {
		return "CAR"
	}
	if b.Type == MOTORCYCLE {
		return "BIKE"
	}
	if b.Type == TRUCK {
		return "TRUCK"
	}
	return ""
}
