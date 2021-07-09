package parking

import (
	"fmt"
	"sort"
)

func CalculateMaxCarsV1(parkingTickets []Ticket) int {
	type tempState struct {
		time    string
		isStart bool
	}

	totalTicketsCount := len(parkingTickets)

	states := make([]tempState, totalTicketsCount*2, totalTicketsCount*2)

	j := 0

	for i :=0; i <len(parkingTickets); i++ {
		states[j] = tempState{parkingTickets[i].Start, true}
		j++
		states[j] = tempState{parkingTickets[i].End, false}
		j++
	}
	sort.Slice(states, func(i, j int) bool {
		firstTime := states[i]
		secondTime := states[j]
		if firstTime.time < secondTime.time {
			return true
		}
		if firstTime.time == secondTime.time {
			return firstTime.isStart
		}
		return false
	})

	maxCarsCount := 0
	counter := 0
	totalArrivalChecked := 0
	for _, state := range states {
		if state.isStart {
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


func CalculateMaxCarsV2(parkingTickets []Ticket) int {
	day := make([]int, 60*24)
	start, end := 0 ,0
	timeMap := getTimeMap()
	for i:=0; i<len(parkingTickets); i++ {
		start = timeMap[parkingTickets[i].Start]
		end = timeMap[parkingTickets[i].End]
		for start <= end {
			day[start]++
			start++
		}
	}
	sort.Ints(day)

	return day[len(day) - 1]
}

func CalculateMaxCarsV3(parkingTickets []Ticket) int {
	timeMap := getTimeMap()
	in := make([]int, 24*60)
	out := make([]int, 24*60)

	start, end := 0,0
	for i := 0; i < len(parkingTickets); i++ {
		start = timeMap[parkingTickets[i].Start]
		end = timeMap[parkingTickets[i].End]
		in[start]++
		out[end]++
	}

	maxCarsCount := 0
	counter := 0
	totalArrivalChecked := 0
	for i := 0; i < len(in); i++ {
		if in[i] > 0 {
			counter += in[i]
			totalArrivalChecked += in[i]
			if maxCarsCount < counter {
				maxCarsCount = counter
			}
			if totalArrivalChecked == len(parkingTickets) {
				break
			}
		}
		if out[i] > 0 {
			counter -= out[i]
		}
	}
	return maxCarsCount
}

func getTimeMap() map[string]int {
	timeMap := make(map[string]int)
	for h:=0; h<24; h++ {
		for m:=0; m<60; m++ {
			timeMap[fmt.Sprintf("%.2d:%.2d", h, m)] = h*60+m
		}
	}
	return timeMap
}