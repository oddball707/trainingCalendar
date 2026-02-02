package model

import (
	"time"
)

// Event is a date entry with the date and description
type Event struct {
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Distance    int       `json:"distance"`
}

type Race struct {
	RaceDate time.Time
	RaceType RaceType
}

type Options struct {
	WeeklyMileage int  `json:"weeklyMileage"`
	BackToBacks   bool `json:"backToBacks"`
	RestDays      int  `json:"restDays"`
	WowIncrease   int  `json:"increase"`
	RestWeekFreq  int  `json:"restWeekFreq"`
	RestWeekLevel int  `json:"restWeekLevel"`
}

type RaceType int

func (r RaceType) GetRaceDistance() int {
	switch r {
	case Half:
		return 13
	case Marathon:
		return 26
	case FiftyK:
		return 31
	case FifyM:
		return 50
	case HundredK:
		return 62
	case HundredM:
		return 100
	}
	return 26
}

const (
	None RaceType = iota
	Half
	Marathon
	FiftyK
	FifyM
	HundredK
	HundredM
	Dynamic
)

func (r RaceType) GetFile() string {
	switch r {
	case Half:
		return "data/half.csv"
	case Marathon:
		return "data/marathon.csv"
	case FiftyK:
		return "data/50k.csv"
	case FifyM:
		return "data/50m.csv"
	case HundredK:
		return "data/100k.csv"
	case HundredM:
		return "data/100m.csv"
	}
	return "data/marathon.csv"
}

func (r RaceType) ToString() string {
	switch r {
	case Half:
		return "Half Marathon"
	case Marathon:
		return "Marathon"
	case FiftyK:
		return "50K Ultra"
	case FifyM:
		return "50 Mile Ultra"
	case HundredK:
		return "100K Ultra"
	case HundredM:
		return "100 Mile Ultra"
	case Dynamic:
		return "Dynamic Schedule"
	}
	return "Error - Reverting to Marathon"
}
