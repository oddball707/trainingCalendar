package model

import (
	"time"
)

//weekStart is the date of Monday
//days is a array of events, starting with monday
type Week struct {
	WeekStart	time.Time   `json:"weekStart"`
	Days		[7]Event	`json:"days"`
}