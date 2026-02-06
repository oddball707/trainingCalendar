package model

import (
	"strconv"
	"strings"
	"time"
)

// Event is a date entry with the date and description
type Event struct {
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Distance    int       `json:"distance"`
}

type RaceType int

const (
	None RaceType = iota
	FiveK
	Half
	Marathon
	FiftyK
	FifyM
	HundredK
	HundredM
	Dynamic
)

type Race struct {
	Name      string
	Type      RaceType
	Distance  float64
	RaceFiles RaceData
	Date      time.Time
}

type RaceData struct {
	TitlesFile       string
	DistancesFile    string
	DescriptionsFile string
}

type DynamicOptions struct {
	WeeklyMileage int     `json:"weeklyMileage"`
	BackToBacks   bool    `json:"backToBacks"`
	RestDays      int     `json:"restDays"`
	WowIncrease   int     `json:"increase"`
	RestWeekFreq  int     `json:"restWeekFreq"`
	RestWeekLevel int     `json:"restWeekLevel"`
	GoalTime      float64 `json:"goalTime"`
}

var RaceTypeMap = map[RaceType]*Race{
	FiveK: {
		Name:     "5k",
		Distance: 3.1,
		RaceFiles: RaceData{
			TitlesFile:       "data/5k.csv",
			DistancesFile:    "data/5k_distance.csv",
			DescriptionsFile: "data/5k_detail.csv",
		},
	},
	Half: {
		Name:     "Half Marathon",
		Distance: 13.1,
		RaceFiles: RaceData{
			TitlesFile: "data/half.csv",
		},
	},
	Marathon: {
		Name:     "Marathon",
		Distance: 26.2,
		RaceFiles: RaceData{
			TitlesFile: "data/marathon.csv",
		},
	},
	FiftyK: {
		Name:     "50K Ultra",
		Distance: 31.1,
		RaceFiles: RaceData{
			TitlesFile: "data/50k.csv",
		},
	},
	FifyM: {
		Name:     "50 Mile Ultra",
		Distance: 50,
		RaceFiles: RaceData{
			TitlesFile: "data/50m.csv",
		},
	},
	HundredK: {
		Name:     "100K Ultra",
		Distance: 62.2,
		RaceFiles: RaceData{
			TitlesFile: "data/100k.csv",
		},
	},
	HundredM: {
		Name:     "100 Mile Ultra",
		Distance: 100,
		RaceFiles: RaceData{
			TitlesFile: "data/100m.csv",
		},
	},
}

func (e *Event) SetDistance() {
	dist, err := strconv.Atoi(strings.Trim(e.Description, " "))
	if strings.ToLower(e.Description) == "rest" ||
		strings.ToLower(e.Description) == "crosstrain" ||
		strings.ToLower(e.Description) == "cross" {
		e.Distance = 0
	} else if err != nil {
		e.Distance = defaultWorkoutDistance
	} else {
		e.Distance = dist
	}
}
