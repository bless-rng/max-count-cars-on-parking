package parking

import (
	"testing"
)

func TestCalculateMaxCarsV1(t *testing.T) {
	parkingTickets := []Ticket{
		{"09:00", "10:08"},
		{"10:20", "11:35"},
		{"12:00", "17:00"},
		{"11:00", "11:30"},
		{"11:20", "12:30"},
		{"11:30", "18:15"},
		{"00:00","00:59"},
		{"00:30","01:00"},
		{"01:00","23:59"},
		{"00:58","02:00"},
	}
	expected := 5


	result := CalculateMaxCarsV1(parkingTickets)
	if result != expected {
		t.Errorf("result is %d; expected %d", result, expected)
	}
}