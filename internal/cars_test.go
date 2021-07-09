package parking

import (
	"testing"
)

func TestCalculateMaxCarsV1(t *testing.T) {
	parkingTickets, expected := getTestData()
	result := CalculateMaxCarsV1(parkingTickets)
	if result != expected {
		t.Errorf("result is %d; expected %d", result, expected)
	}
}

func TestCalculateCarsByArray(t *testing.T) {
	parkingTickets, expected := getTestData()
	result := CalculateCarsByArray(parkingTickets)
	if result != expected {
		t.Errorf("result is %d; expected %d", result, expected)
	}
}

func BenchmarkCalculateMaxCarsV1(b *testing.B) {
	parkingTickets, _ := getTestData()
	for i := 0; i < b.N; i++ {
		CalculateMaxCarsV1(parkingTickets)
	}
}

func BenchmarkCalculateMaxCarsByArray(b *testing.B) {
	parkingTickets, _ := getTestData()
	for i := 0; i < b.N; i++ {
		CalculateCarsByArray(parkingTickets)
	}
}

func getTestData() (parkingTickets []Ticket, expected int) {
	tickets, expected := []Ticket{
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
	}, 5
	repeatTimes := 15
	var resultTickets []Ticket
	for i := 0; i < repeatTimes; i++ {
		resultTickets = append(resultTickets, tickets...)
	}
	expected = expected*repeatTimes
	return resultTickets, expected
}