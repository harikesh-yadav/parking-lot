package parkinglot

import (
	"fmt"
	"time"
)

type Ticket struct {
	Level   Level
	Spot    ParkingSpot
	Vehicle Vehicle
}

func NewTicket(l Level, s ParkingSpot, v Vehicle) *Ticket {
	return &Ticket{
		Level:   l,
		Spot:    s,
		Vehicle: v,
	}
}

func (t *Ticket) Print() {
	fmt.Println("Floor:", t.Level.Floor)
	fmt.Println("Spot:", t.Spot.SpotNumber)
	fmt.Println("Vehicle Type:", t.Vehicle.VehicleName())
	fmt.Println("Number Plate", t.Vehicle.GetLicencePlate())
	fmt.Println("Park Time:", time.Now())
	fmt.Println("-----------------------------------")
}
