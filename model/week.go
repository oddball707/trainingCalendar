package model

import (
	"fmt"
	"strconv"
	"strings"
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

const defaultWorkoutDistance = 6

func (w *Week) SetDistance() {
	actualMileage := 0
	for _, day := range w.Days {
		fmt.Println(day.Description)
		mile, err := strconv.Atoi(strings.Trim(day.Description, " "))
		if strings.ToLower(day.Description) == "rest" || strings.ToLower(day.Description) == "crosstrain" || strings.ToLower(day.Description) == "cross" {
			continue
		} else if err != nil {
			actualMileage += defaultWorkoutDistance
		} else {
			actualMileage += mile
		}
	}
	fmt.Printf("Actual Weekly Mileage: %v\n", actualMileage)
	w.TotalDistance = actualMileage
}
