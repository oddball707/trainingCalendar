package service

import (
	"errors"
	"fmt"
	"log"
	"math"
	"regexp"
	"time"

	m "github.com/oddball707/trainingCalendar/model"
)

// Assumes 3 easy + 1 tempo + 1 med + 1 long = 110%
const (
	MinimumWeeklyMileage = 15
	EasyDayPercentage    = 0.1
	MedDayPercentage     = 0.17
	LongDayPercentage    = 0.4
	TempoDayPercentage   = 0.23
)

type Generator struct {
	Race                 *m.Race
	CurrentWeeklyMileage int
	RestDays             int
	BackToBacks          bool
	RestWeekFreq         int
	RestWeekLevel        float64
}

type ScheduleGenerator interface {
	CreateScheduleForRace(r *m.Race) (m.Schedule, error)
	CreateScheduleStartingNow(totalLength int) (m.Schedule, error)
}

func NewGenerator(options *m.Options, race *m.Race) *Generator {
	fmt.Printf("Submitted Options: %v\n", options)
	if options.RestWeekFreq == 0 {
		options.RestWeekFreq = 4
	}
	if options.RestWeekLevel == 0 {
		options.RestWeekLevel = 60
	}
	return &Generator{
		Race:                 race,
		CurrentWeeklyMileage: options.WeeklyMileage,
		RestDays:             options.RestDays,
		BackToBacks:          options.BackToBacks,
		RestWeekFreq:         options.RestWeekFreq,
		RestWeekLevel:        float64(options.RestWeekLevel) / 100,
	}
}

func (g *Generator) CreateScheduleForRace() (m.Schedule, error) {
	if g.RestDays > 4 {
		return nil, errors.New("too many rest days")
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
	lastMonday := PrevMonday(g.Race.RaceDate)

	totalLength := int(lastMonday.Sub(firstMonday).Hours() / 24 / 7)

	fmt.Printf("Total hours in schedule: %v\n", int(lastMonday.Sub(firstMonday).Hours()))
	fmt.Printf("Total weeks in schedule: %v\n", totalLength)
	fmt.Printf("Last week starts on %v\n", lastMonday)
	var sched m.Schedule
	for firstMonday.Before(g.Race.RaceDate) {
		fmt.Printf("Week # %v\n", weekNumber)
		var week *m.Week
		if weekNumber%g.RestWeekFreq == 0 {
			fmt.Printf("Target Rest Weekly Mileage: %v\n", weeklyMileage*(float64(g.RestWeekLevel)))
			week = g.generateWeek(firstMonday, weeklyMileage*(float64(g.RestWeekLevel)))
		} else {
			fmt.Printf("Target Weekly Mileage: %v\n", weeklyMileage)
			week = g.generateWeek(firstMonday, weeklyMileage)
		}
		week.WowIncrease = Increase(sched, week.TotalDistance, g.RestWeekLevel)
		weeklyMileage += weeklyGrowth(weeklyMileage)

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

func Increase(sched m.Schedule, actualMileage int, restWeekLevel float64) string {
	if len(sched) < 1 {
		return "-"
	}
	wow := 1 - (float64(sched[len(sched)-1].TotalDistance) / float64(actualMileage))
	if wow < 0 {
		return "-"
	} else if wow > (1 - restWeekLevel) {
		wow = 1 - (float64(sched[len(sched)-2].TotalDistance) / float64(actualMileage))
	}
	r, _ := regexp.Compile(`\.?0*$`)
	return fmt.Sprintf("%s%%\n", r.ReplaceAllString(fmt.Sprintf("%.2f", 100*wow), ""))
}

func (g *Generator) setTaper(sched m.Schedule, weeklyMileage float64) m.Schedule {
	secondlastWeek := sched[len(sched)-2]
	lastMonday := sched[len(sched)-1].WeekStart
	taper := g.generateWeek(secondlastWeek.WeekStart, weeklyMileage/3)
	taper.SetDistance()
	taper.WowIncrease = "-"

	monday := m.Event{Date: lastMonday, Title: "Rest", Distance: 0}
	tuesday := m.Event{Date: lastMonday.AddDate(0, 0, 1), Title: "2", Distance: 2}
	wednesday := m.Event{Date: lastMonday.AddDate(0, 0, 2), Title: "3", Distance: 3}
	thursday := m.Event{Date: lastMonday.AddDate(0, 0, 3), Title: "Rest", Distance: 0}
	friday := m.Event{Date: lastMonday.AddDate(0, 0, 4), Title: "2", Distance: 2}
	saturday := m.Event{Date: lastMonday.AddDate(0, 0, 5), Title: "Race Day!", Distance: g.Race.RaceType.GetRaceDistance()}
	sunday := m.Event{Date: lastMonday.AddDate(0, 0, 6), Title: "Rest", Distance: 0}

	raceWeek := &m.Week{
		WeekStart:     lastMonday,
		Days:          [7]m.Event{monday, tuesday, wednesday, thursday, friday, saturday, sunday},
		TotalDistance: 7,
		WowIncrease:   "--race week--",
	}
	sched[len(sched)-2] = taper
	sched[len(sched)-1] = raceWeek
	return sched
}

func weeklyGrowth(x float64) float64 {
	return math.Log(x + 2)
}

func (g *Generator) generateWeek(firstMonday time.Time, weeklyMileage float64) *m.Week {
	longDay := int(math.Floor(weeklyMileage * LongDayPercentage))
	easyDay := int(math.Ceil(weeklyMileage * EasyDayPercentage))
	medDay := int(math.Floor(weeklyMileage * MedDayPercentage))
	tempoDay := int(math.Ceil(weeklyMileage * TempoDayPercentage))

	//[Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	days := [7]int{0, 0, 0, 0, 0, longDay, 0}

	switch g.RestDays {
	case 4:
		//3 days = .4+.4+.23 = 103%
		if g.BackToBacks {
			days[2] = longDay - 1
			days[6] = tempoDay
		} else {
			days[1] = longDay - 1
			days[3] = tempoDay
		}
	case 3:
		//4 days .4+.23+.23+.17= 103%
		if g.BackToBacks {
			days[6] = tempoDay
		} else {
			days[0] = tempoDay
		}
		days[1] = medDay
		days[3] = tempoDay - 1
	case 2:
		//5 days .1+.1+.17+.23+.4 = 100%
		if g.BackToBacks {
			days[1] = easyDay
			days[2] = medDay
			days[3] = easyDay
			days[6] = tempoDay
		} else {
			days[1] = tempoDay
			days[2] = easyDay
			days[3] = medDay
			days[6] = easyDay
		}
	case 1:
		//.1+.1+.1+.17+.17+.4 = 104
		if g.BackToBacks {
			days[1] = easyDay
			days[2] = medDay
			days[3] = easyDay
			days[4] = easyDay - 1
			days[6] = medDay
		} else {
			days[1] = medDay
			days[2] = easyDay
			days[3] = medDay
			days[4] = easyDay - 1
			days[6] = easyDay
		}
	case 0:
		//.1+.1+.1+.1+.1+.17+.4=107%
		if g.BackToBacks {
			days[0] = easyDay - 1
			days[1] = easyDay
			days[2] = easyDay - 1
			days[3] = easyDay
			days[4] = easyDay - 1
			days[6] = medDay
		} else {
			days[0] = easyDay - 1
			days[1] = easyDay
			days[2] = medDay
			days[3] = easyDay
			days[4] = easyDay - 1
			days[6] = easyDay - 1
		}
	}

	actualMileage := 0
	weekDays := [7]m.Event{}
	for i, day := range days {
		actualMileage += day
		dailyEvent := m.Event{Date: firstMonday.AddDate(0, 0, i), Title: distToStr(days[i]), Distance: days[i]}
		weekDays[i] = dailyEvent
	}
	fmt.Printf("Actual Weekly Mileage: %v\n", actualMileage)

	return &m.Week{
		WeekStart:     firstMonday,
		Days:          weekDays,
		TotalDistance: actualMileage,
	}
}

func distToStr(dist int) string {
	if dist == 0 {
		return "Rest"
	}
	return fmt.Sprintf("%v", dist)
}
