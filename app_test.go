package main

import "testing"

func TestCalculateMaxCarsV1(t *testing.T) {
	parkingTickets := []ParkingTicket{
		{"09:00", "10:08"},
		{"10:20", "11:35"},
		{"12:00", "17:00"},
		{"11:00", "11:30"},
		{"11:20", "12:30"},
		{"11:30", "18:15"},
	}
	expected := 4


	result := calculateMaxCarsV1(parkingTickets)
	if result != expected {
		t.Errorf("result is %d; expected %d", result, expected)
	}
}