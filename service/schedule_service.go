package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"time"

	m "trainingCalendar/model"

	"github.com/jordic/goics"
)

type Service struct {
	//TODO: dependency injection for data
}

type ScheduleService interface {
	GetSchedule(r *m.Race) (*m.Schedule, error)
	CreateIcal(r *m.Race) (*os.File, error)
	LoadCalendar(r *m.Race) (m.Schedule, error)
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetSchedule(r *m.Race) (*m.Schedule, error) {
	schedule, err := s.LoadCalendar(r)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (s *Service) CreateIcal(r *m.Race) (*os.File, error) {
	weeks, err := s.LoadCalendar(r)
	if err != nil {
		return nil, err
	}
	b := &bytes.Buffer{}
	enc := goics.NewICalEncode(b)
	enc.Encode(weeks)

	calFile := "out/training.ics"
	if _, err := os.Stat(calFile); os.IsNotExist(err) {
		os.MkdirAll("out", 0700)
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

func (s *Service) LoadCalendar(race *m.Race) (m.Schedule, error) {
	lines, err := s.readRaceFile(race)
	if err != nil {
		return nil, err
	}

	monBeforeRace := prevMonday(race.RaceDate)
	firstMonday := monBeforeRace.AddDate(0, 0, (len(lines)-1)*-7)
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

func prevMonday(day time.Time) time.Time {
	if day.Weekday() == time.Sunday {
		return day.AddDate(0, 0, -6)
	} else {
		return day.AddDate(0, 0, int(day.Weekday())-1)
	}
}
