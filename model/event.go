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
	Half RaceType = iota
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
