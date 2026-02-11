package model

import (
	"log"
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

const defaultWorkoutDistance = 6

func (w *Week) SetDistance() {
	actualMileage := 0
	for _, day := range w.Days {
		log.Println("Title: " + day.Title)
		log.Println("Description: " + day.Description)
		log.Println("Distance: " + strconv.Itoa(day.Distance))

		actualMileage += day.Distance
	}
	log.Printf("Actual Weekly Mileage: %v", actualMileage)
	w.TotalDistance = actualMileage
}
