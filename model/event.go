package model

import (
	"time"
)

// Event is a date entry with the date and description
type Event struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type Race struct {
	RaceDate time.Time
	RaceType RaceType
}

type RaceType int

const (
	None RaceType = iota
	Half
	Marathon
	FiftyK
	FifyM
	HundredK
	HundredM
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
	}
	return "Error - Reverting to Marathon"
}
