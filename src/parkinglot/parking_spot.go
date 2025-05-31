package parkinglot

type ParkingSpot struct {
	SpotNumber    int
	VehicleType   VehicleType
	ParkedVehical Vehicle
}

func NewParkingSpot(sn int, vt VehicleType) *ParkingSpot {
	return &ParkingSpot{
		SpotNumber:  sn,
		VehicleType: vt,
	}
}

func (p *ParkingSpot) IsAvailable() bool {
	return p.ParkedVehical == nil
}
func (p *ParkingSpot) ParkVehicle(vehicle Vehicle) bool {
	if p.IsAvailable() && p.VehicleType == vehicle.GetType() {
		p.ParkedVehical = vehicle
		return true
	}
	return false
}

func (p *ParkingSpot) GetSpotNumber() int {
	return p.SpotNumber
}
func (p ParkingSpot) GetVehicleType() VehicleType {
	return p.VehicleType
}

func (p *ParkingSpot) GetVehicle() Vehicle {
	return p.ParkedVehical
}
