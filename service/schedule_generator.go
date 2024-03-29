package service

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
	"strconv"

	m "github.com/oddball707/trainingCalendar/model"
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
	firstMonday := NextMonday(time.Now())
	lastMonday := PrevMonday(r.RaceDate)

	totalLength := int(lastMonday.Sub(firstMonday).Hours()/24/7)

	fmt.Printf("Total hours in schedule: %v\n", int(lastMonday.Sub(firstMonday).Hours()))
	fmt.Printf("Total weeks in schedule: %v\n", totalLength)
	fmt.Printf("Last week starts on %v\n", lastMonday)
	var sched m.Schedule
	for firstMonday.Before(r.RaceDate) {
		fmt.Printf("Week # %v\n", weekNumber)
		var week *m.Week
		if weekNumber % 3 == 0 {
			fmt.Printf("Weekly Mileage: %v\n", weeklyMileage / 2)
			week = g.generateWeek(firstMonday, weeklyMileage / 2)
		} else {
			fmt.Printf("Target Weekly Mileage: %v\n", weeklyMileage)
			week = g.generateWeek(firstMonday, weeklyMileage)
			actualMileage := 0
			for _, day := range week.Days {
				mile, _ := strconv.Atoi(day.Description)
				actualMileage += mile
			}
			fmt.Printf("Actual Weekly Mileage: %v\n", actualMileage)
			weeklyMileage += weeklyGrowth(weeklyMileage)
		}
		sched = append(sched, week)
		firstMonday = firstMonday.AddDate(0, 0, 7)
		weekNumber += 1
	}

	sched = g.setTaper(sched, weeklyMileage)

	return sched, nil
}

func CreateScheduleStartingNow(totalLength int) (schedule *m.Schedule, err error) {
	return nil, nil
}

func (g *Generator) setTaper(sched m.Schedule, weeklyMileage float64) m.Schedule {
	secondlastWeek := sched[len(sched)-2]
	lastMonday := sched[len(sched)-1].WeekStart
	taper := g.generateWeek(secondlastWeek.WeekStart, weeklyMileage / 3)
	
	monday := m.Event{lastMonday, "Rest"}
	tuesday := m.Event{lastMonday.AddDate(0, 0, 1), "2"}
	wednesday := m.Event{lastMonday.AddDate(0, 0, 2), "3"}
	thursday := m.Event{lastMonday.AddDate(0, 0, 3), "Rest"}
	friday := m.Event{lastMonday.AddDate(0, 0, 4), "2"}
	saturday := m.Event{lastMonday.AddDate(0, 0, 5), "Race Day!"}
	sunday := m.Event{lastMonday.AddDate(0, 0, 6),  "Rest"}

	raceWeek := &m.Week{
		WeekStart: lastMonday,
		Days: [7]m.Event{monday, tuesday, wednesday, thursday, friday, saturday, sunday},
	}
	sched[len(sched)-2] = taper
	sched[len(sched)-1] = raceWeek
	return sched
}

func weeklyGrowth(x float64) float64 {
	return math.Log(x + 2)
}

func (g *Generator) generateWeek(firstMonday time.Time, weeklyMileage float64) *m.Week {
	longDay := fmt.Sprintf("%v", int(math.Floor(weeklyMileage * LongDayPercentage)))
	easyDay := fmt.Sprintf("%v", int(math.Ceil(weeklyMileage * EasyDayPercentage)))
	medDay := fmt.Sprintf("%v", int(math.Floor(weeklyMileage * MedDayPercentage)))
	tempoDay := fmt.Sprintf("%v", int(math.Ceil(weeklyMileage * TempoDayPercentage)))

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
			wednesdayMileage = medDay
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

	return &m.Week{
		WeekStart: firstMonday,
		Days: [7]m.Event{monday, tuesday, wednesday, thursday, friday, saturday, sunday},
	}
}