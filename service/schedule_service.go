package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"time"

	m "github.com/oddball707/trainingCalendar/model"

	"github.com/jordic/goics"
)

type Service struct {
	//TODO: dependency injection for data
}

type ScheduleService interface {
	GetSchedule(r *m.Race, o *m.Options) (*m.Schedule, error)
	CreateIcal(r *m.Race, o *m.Options) (*os.File, error)
	LoadCalendar(r *m.Race, o *m.Options) (m.Schedule, error)
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetSchedule(r *m.Race, o *m.Options) (*m.Schedule, error) {
	log.Printf("Creating schedule for a %s that starts on %s", r.RaceType.ToString(), r.RaceDate.Format(m.DateLayout))
	schedule, err := s.LoadCalendar(r, o)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (s *Service) CreateIcal(r *m.Race, o *m.Options) (*os.File, error) {
	log.Printf("Creating an ical for a %s that starts on %s", r.RaceType.ToString(), r.RaceDate.Format(m.DateLayout))
	weeks, err := s.LoadCalendar(r, o)
	if err != nil {
		return nil, err
	}
	weeks.Print()
	b := &bytes.Buffer{}
	enc := goics.NewICalEncode(b)
	enc.Encode(weeks)

	calFile := "out/training.ics"
	err = os.MkdirAll("out", 0700)
	if err != nil {
		log.Print("Error creating out dir: " + err.Error())
	}
	f, err := os.Create(calFile)
	if err != nil {
		log.Print("Error creating ical: " + err.Error())
		return nil, err
	}

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

func (s *Service) LoadCalendar(race *m.Race, options *m.Options) (m.Schedule, error) {
	if race.RaceType == m.Dynamic {
		return generateSchedule(race, options)
	}
	lines, err := s.readRaceFile(race)
	if err != nil {
		return nil, err
	}

	firstMonday := s.startDate(race.RaceDate, len(lines))
	var sched m.Schedule

	// Loop through lines & turn into object
	for _, line := range lines {
		var days [7]m.Event
		for i, desc := range line {
			days[i] = m.Event{firstMonday.AddDate(0, 0, i), desc}
		}
		wk := &m.Week{
			WeekStart: firstMonday,
			Days:      days,
		}
		sched = append(sched, wk)
		firstMonday = firstMonday.AddDate(0, 0, 7)
	}

	return sched, nil
}

func (s *Service) readRaceFile(r *m.Race) ([][]string, error) {
	f, err := os.Open(r.RaceType.GetFile())
	if err != nil {
		log.Print("Failed to open csv file: " + r.RaceType.GetFile())
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

func generateSchedule(race *m.Race, o *m.Options) (schedule m.Schedule, err error) {
	generator := NewGenerator(o.WeeklyMileage, o.RestDays, o.BackToBacks)
	return generator.CreateScheduleForRace(race)
}
