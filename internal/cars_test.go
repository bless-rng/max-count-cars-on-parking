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
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
		{"00:00", "23:59"},
	}, 10
	repeatTimes := 1000
	var resultTickets []Ticket
	for i := 0; i < repeatTimes; i++ {
		resultTickets = append(resultTickets, tickets...)
	}
	expected = expected*repeatTimes
	return resultTickets, expected
}