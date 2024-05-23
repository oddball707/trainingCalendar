package model

import (
	"fmt"
	"strconv"
	"time"
)

// weekStart is the date of Monday
// days is a array of events, starting with monday
// TotalDistance is the total mileage of this week
// WowIncrease is the increase in mileage over the previous week, as a percentage
type Week struct {
	WeekStart     time.Time `json:"weekStart"`
	Days          [7]Event  `json:"days"`
	TotalDistance int       `json:"totalDistance"`
	WowIncrease   string    `json:"wowIncrease"`
}

func (w *Week) SetDistance() {
	actualMileage := 0
	for _, day := range w.Days {
		mile, _ := strconv.Atoi(day.Description)
		actualMileage += mile
	}
	fmt.Printf("Actual Weekly Mileage: %v\n", actualMileage)
	w.TotalDistance = actualMileage
}
