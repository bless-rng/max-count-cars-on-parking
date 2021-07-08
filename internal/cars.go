package parking

import "sort"

func CalculateMaxCarsV1(parkingTickets []Ticket) int {
	type tempState struct {
		time    string
		isStart bool
	}

	totalTicketsCount := len(parkingTickets)

	var states []tempState

	for _, parking := range parkingTickets {
		states = append(states, tempState{
			parking.Start,
			true,
		})
		states = append(states, tempState{
			parking.End,
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
			if maxCarsCount < counter {
				maxCarsCount = counter
			}
			if totalArrivalChecked == totalTicketsCount {
				break
			}
		} else {
			counter -= 1
		}
	}
	return maxCarsCount
}