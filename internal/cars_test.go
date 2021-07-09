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

func TestCalculateMaxCarsV2(t *testing.T) {
	parkingTickets, expected := getTestData()
	result := CalculateMaxCarsV2(parkingTickets)
	if result != expected {
		t.Errorf("result is %d; expected %d", result, expected)
	}
}

func TestCalculateMaxCarsV3(t *testing.T) {
	parkingTickets, expected := getTestData()
	result := CalculateMaxCarsV3(parkingTickets)
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

func BenchmarkCalculateMaxCarsV2(b *testing.B) {
	parkingTickets, _ := getTestData()
	for i := 0; i < b.N; i++ {
		CalculateMaxCarsV2(parkingTickets)
	}
}

func BenchmarkCalculateMaxCarsV3(b *testing.B) {
	parkingTickets, _ := getTestData()
	for i := 0; i < b.N; i++ {
		CalculateMaxCarsV3(parkingTickets)
	}
}

func getTestData() (parkingTickets []Ticket, expected int) {
	tickets, expected := []Ticket{
		{"00:00", "01:00"},
		{"00:00", "02:00"},
		{"03:00", "04:00"},
		{"05:00", "06:00"},
		{"07:00", "08:00"},
		{"09:00", "10:00"},
		{"11:00", "12:00"},
		{"13:00", "14:00"},
		{"15:00", "16:00"},
		{"17:00", "18:00"},
	}, 2
	repeatTimes := 1000
	var resultTickets []Ticket
	for i := 0; i < repeatTimes; i++ {
		resultTickets = append(resultTickets, tickets...)
	}
	expected = expected*repeatTimes
	return resultTickets, expected
}