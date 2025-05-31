package parkinglot

import "fmt"

type Level struct {
	Floor        int
	ParkingSpots []*ParkingSpot
}

func NewLevel(floor int, numSpot int) *Level {
	level := &Level{Floor: floor}
	bikeSpot := int(float64(numSpot) * 0.50)
	carSpot := int(float64(numSpot) * 0.40)
	for i := 1; i <= bikeSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, NewParkingSpot(i, MOTORCYCLE))
	}
	for i := bikeSpot + 1; i <= bikeSpot+bikeSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, NewParkingSpot(i, CAR))
	}
	for i := bikeSpot + carSpot + 1; i <= numSpot; i++ {
		level.ParkingSpots = append(level.ParkingSpots, NewParkingSpot(i, TRUCK))
	}
	return level
}

func (l *Level) ParkedVehicle(v Vehicle) bool {
	for i, spot := range l.ParkingSpots {
		if spot.IsAvailable() && spot.VehicleType == v.GetType() {
			spot.ParkVehicle(v)
			t := NewTicket(*l, *l.ParkingSpots[i], v)
			t.Print()
			return true
		}
	}
	return false
}

func (l *Level) UnParkedVehicle(v Vehicle) bool {
	for _, spot := range l.ParkingSpots {
		if !spot.IsAvailable() && spot.GetVehicle() == v {
			spot.ParkedVehical = nil
			return true
		}
	}
	return false
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.ParkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
		}
		fmt.Println("Floor:", l.Floor, ", Spot Number:", spot.GetSpotNumber(), ", Status:", status)
	}
}
