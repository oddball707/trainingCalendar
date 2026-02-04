package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jordic/goics"

	m "github.com/oddball707/trainingCalendar/model"
)

type Service struct {
	//TODO: dependency injection for data
}

type ScheduleService interface {
	GetSchedule(r *m.Race, o *m.DynamicOptions) (*m.Schedule, error)
	CreateIcal(r *m.Race, o *m.DynamicOptions, f string) (*os.File, error)
	LoadCalendar(r *m.Race, o *m.DynamicOptions) (m.Schedule, error)
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetSchedule(r *m.Race, o *m.DynamicOptions) (*m.Schedule, error) {
	log.Printf("Creating schedule for a %s that starts on %s", r.Name, r.Date.Format(m.DateLayout))
	schedule, err := s.LoadCalendar(r, o)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (s *Service) CreateIcal(r *m.Race, o *m.DynamicOptions, filePath string) (*os.File, error) {
	log.Printf("Creating an ical for a %s race on %s", r.Name, r.Date.Format(m.DateLayout))
	weeks, err := s.LoadCalendar(r, o)
	if err != nil {
		return nil, err
	}
	weeks.Print()
	b := &bytes.Buffer{}
	enc := goics.NewICalEncode(b)
	enc.Encode(weeks)

	calFile := filePath + "/training.ics"
	// err = os.MkdirAll("out", 0700)
	// if err != nil {
	// 	log.Print("Error creating out dir: " + err.Error())
	// }
	f, err := os.Create(calFile)
	if err != nil {
		log.Print("Error creating ical file: " + err.Error())
		return nil, err
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	b.WriteTo(w)
	if err != nil {
		log.Print("Error writing ical: " + err.Error())
		return nil, err
	}

	w.Flush()
	if err != nil {
		log.Print("Error flushing writer: " + err.Error())
		return nil, err
	}

	return f, nil
}

func (s *Service) LoadCalendar(race *m.Race, options *m.DynamicOptions) (m.Schedule, error) {
	if race.Type == m.Dynamic {
		return generateSchedule(race, options)
	}

	mainLines, err := s.readRaceFile(race.RaceFiles.TitlesFile)
	if err != nil {
		return nil, err
	}

	descLines, err := s.readRaceFile(race.RaceFiles.DescriptionsFile)
	if err != nil {
		return nil, err
	}

	distLines, err := s.readRaceFile(race.RaceFiles.DistancesFile)
	if err != nil {
		return nil, err
	} else if distLines == nil {
		log.Print("No distances file, using main file for distances")
		distLines = mainLines
	}

	firstMonday := s.startDate(race.Date, len(mainLines))
	var sched m.Schedule

	// Loop through lines & turn into object
	for week, line := range mainLines {
		var days [7]m.Event
		var weeksDescriptions []string

		weeksDistances := distLines[week]
		fmt.Printf("Week %d distances: %v\n", week, weeksDistances)
		if descLines != nil {
			weeksDescriptions = descLines[week]
		}
		for day, title := range line {
			dist, err := strconv.Atoi(weeksDistances[day])
			if err != nil {
				log.Print("Error converting calendar entry to distance: " + weeksDistances[day] + "; " + err.Error())
				return nil, err
			}
			days[day] = m.Event{Date: firstMonday.AddDate(0, 0, day), Title: title, Distance: dist}

			if weeksDescriptions != nil && weeksDescriptions[day] != "" {
				days[day].Description = weeksDescriptions[day]
			}
		}

		wk := &m.Week{
			WeekStart: firstMonday,
			Days:      days,
		}
		wk.SetDistance()
		wk.WowIncrease = Increase(sched, wk.TotalDistance, 0.7)
		sched = append(sched, wk)
		firstMonday = firstMonday.AddDate(0, 0, 7)
	}

	return sched, nil
}

func (s *Service) readRaceFile(path string) ([][]string, error) {
	if path == "" {
		return nil, nil
	}

	f, err := os.Open(path)
	if err != nil {
		log.Print("Failed to open csv file: " + path)
		return nil, err
	}
	defer f.Close()

	// Read File into a Variable
	f.Seek(0, 0)
	reader := csv.NewReader(f)
	if _, err := reader.Read(); err != nil { //read header
		log.Print("Error reading csv header: " + err.Error())
		return nil, err
	}

	lines, err := reader.ReadAll()
	if err != nil {
		log.Print("Error Reading Csv: " + err.Error())
		return nil, err
	}

	return lines, nil
}

func (s *Service) startDate(raceDate time.Time, weeksInSched int) time.Time {
	monBeforeRace := PrevMonday(raceDate)
	return monBeforeRace.AddDate(0, 0, (weeksInSched-1)*-7)
}

func generateSchedule(race *m.Race, o *m.DynamicOptions) (schedule m.Schedule, err error) {
	generator := NewGenerator(o, race)
	return generator.CreateScheduleForRace()
}
