package main

import (
	parking "parking/internal"
)

func main() {
	parkingTickets := []parking.Ticket{
		{"09:00", "10:08"},
		{"10:20", "11:35"},
		{"12:00", "17:00"},
		{"11:00", "11:30"},
		{"11:20", "12:30"},
		{"11:30", "18:15"},
	}

	maxCarsCount := 0
	if len(parkingTickets) >= 150 {
		maxCarsCount = parking.CalculateCarsByArray(parkingTickets)
	} else {
		maxCarsCount = parking.CalculateMaxCarsV1(parkingTickets)
	}

	println("Max count of cars on ParkingTicket is ", maxCarsCount)
}