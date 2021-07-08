package main

import "sort"

func main() {
	parkingTickets := []ParkingTicket{
		{"09:00", "10:08"},
		{"10:20", "11:35"},
		{"12:00", "17:00"},
		{"11:00", "11:30"},
		{"11:20", "12:30"},
		{"11:30", "18:15"},
	}

	maxCarsCount := calculateMaxCarsV1(parkingTickets)

	println("Max count of cars on ParkingTicket is ", maxCarsCount)
}

type tempState struct {
	time    string
	isStart bool
}

type ParkingTicket struct {
	start string
	end   string
}

func calculateMaxCarsV1(parkingTickets []ParkingTicket) int {
	totalTicketsCount := len(parkingTickets)

	var states []tempState

	for _, parking := range parkingTickets {
		states = append(states, tempState{
			parking.start,
			true,
		})
		states = append(states, tempState{
			parking.end,
			false,
		})
	}
	sort.Slice(states, func(i, j int) bool {
		firstTime := states[i]
		secondTime := states[j]
		if firstTime.time < secondTime.time {
			return true
		}
		if firstTime.time == secondTime.time {
			return firstTime.isStart == true
		}
		return false
	})

	maxCarsCount := 0
	counter := 0
	totalArrivalChecked := 0
	for _, time := range states {
		if time.isStart {
			counter += 1
			totalArrivalChecked += 1
		} else {
			counter -= 1
		}
		if maxCarsCount < counter {
			maxCarsCount = counter
		}
		if totalArrivalChecked == totalTicketsCount {
			break
		}
	}
	return maxCarsCount
}
