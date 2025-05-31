package parkinglot

import "sync"

type ParkingLot struct {
	Levels []*Level
}

var instance *ParkingLot
var once sync.Once

func GetParkingLotInstance() *ParkingLot {
	once.Do(func() {
		instance = &ParkingLot{
			Levels: make([]*Level, 0),
		}
	})
	return instance
}

func (p *ParkingLot) AddLevel(l *Level) {
	p.Levels = append(p.Levels, l)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.Levels {
		if level.ParkedVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.Levels {
		if level.UnParkedVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.Levels {
		level.DisplayAvailability()
	}
}
