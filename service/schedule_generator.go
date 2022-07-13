package service

import (
	"errors"
	"fmt"
	"log"
	"math"

	m "trainingCalendar/model"
)

const (
	MinimumWeeklyMileage = 15
	EasyDayPercentage = 0.1
	MedDayPercentage = 0.17
	LongDayPercentage = 0.4
	TempoDayPercentage = 0.23
)

type Generator struct {
	CurrentWeeklyMileage int
	RestDays			 int
	BackToBacks			 bool
}

type ScheduleGenerator interface {
	CreateScheduleForRace(r *m.Race) (m.Schedule, error)
	CreateScheduleStartingNow(totalLength int) (m.Schedule, error)
}

func NewGenerator(currentMileage int, desiredRestDays int, backToBacks bool) *Generator {
	return &Generator{
		CurrentWeeklyMileage: currentMileage,
		RestDays: desiredRestDays,
		BackToBacks: backToBacks,
	}
}

func (g *Generator) CreateScheduleForRace(r *m.Race) (m.Schedule, error) {
	if g.RestDays > 4 {
		return nil, errors.New("Too many rest days")
	} else if g.RestDays > 3 {
		log.Print("Suggest at most 3 rest days per week")
	} else if g.RestDays < 1 {
		log.Print("Suggest at least 1 rest day per week")
	}

	weeklyMileage := float64(g.CurrentWeeklyMileage)
	if g.CurrentWeeklyMileage < MinimumWeeklyMileage {
		weeklyMileage = float64(MinimumWeeklyMileage)
	}

	weekNumber := 1
	firstMonday := NextMonday()
	lastMonday := PrevMonday(r.RaceDate)

	totalLength := int(lastMonday.Sub(firstMonday).Hours()/24/7)

	var sched m.Schedule
	for weekNumber < totalLength {
		longDay := fmt.Sprintf("%f", math.Floor(weeklyMileage * LongDayPercentage))
		easyDay := fmt.Sprintf("%f", math.Ceil(weeklyMileage * EasyDayPercentage))
		medDay := fmt.Sprintf("%f", math.Floor(weeklyMileage * MedDayPercentage))
		tempoDay := fmt.Sprintf("%f", math.Ceil(weeklyMileage * TempoDayPercentage))

		var (
			sundayMileage string = "Rest"
			mondayMileage string = "Rest"
			tuesdayMileage string = "Rest"
			wednesdayMileage string = "Rest"
			thursdayMileage string = "Rest"
			fridayMileage string = "Rest"
		)
		

		switch g.RestDays {
		case 4:
			tuesdayMileage = longDay
			thursdayMileage = tempoDay
		case 3:
			mondayMileage = tempoDay
			tuesdayMileage = medDay
			thursdayMileage = tempoDay
		case 2:
			tuesdayMileage = easyDay
			thursdayMileage = easyDay
			if g.BackToBacks {
				sundayMileage = tempoDay
				wednesdayMileage = easyDay
			} else {
				sundayMileage = easyDay
				wednesdayMileage = medDay
			}
		case 1:
			if g.BackToBacks {
				tuesdayMileage = easyDay
				wednesdayMileage = medDay
				thursdayMileage = easyDay
				fridayMileage = easyDay
				sundayMileage = tempoDay
			} else {			
				tuesdayMileage = medDay
				wednesdayMileage = easyDay
				thursdayMileage = tempoDay
				fridayMileage = easyDay
				sundayMileage = easyDay
			}
		case 0:
			if g.BackToBacks {
				mondayMileage = easyDay
				tuesdayMileage = easyDay
				wednesdayMileage = medDay
				thursdayMileage = easyDay
				fridayMileage = easyDay
				sundayMileage = tempoDay
			} else {			
				tuesdayMileage = medDay
				wednesdayMileage = easyDay
				thursdayMileage = tempoDay
				fridayMileage = easyDay
				sundayMileage = easyDay
			}
		}

		monday := m.Event{firstMonday, mondayMileage}
		tuesday := m.Event{firstMonday.AddDate(0, 0, 1), tuesdayMileage}
		wednesday := m.Event{firstMonday.AddDate(0, 0, 2), wednesdayMileage}
		thursday := m.Event{firstMonday.AddDate(0, 0, 3), thursdayMileage}
		friday := m.Event{firstMonday.AddDate(0, 0, 4), fridayMileage}
		saturday := m.Event{firstMonday.AddDate(0, 0, 5), longDay}
		sunday := m.Event{firstMonday.AddDate(0, 0, 6),  sundayMileage}

		week := &m.Week{
			WeekStart: firstMonday,
			Days: [7]m.Event{monday, tuesday, wednesday, thursday, friday, saturday, sunday},
		}
		sched = append(sched, week)
		firstMonday = firstMonday.AddDate(0, 0, 7)
		weekNumber += 1
	}

	return sched, nil
}

func CreateScheduleStartingNow(totalLength int) (schedule *m.Schedule, err error) {
	return nil, nil
}