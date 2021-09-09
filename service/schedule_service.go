package service

import (
	"bufio"
	"bytes"
	"errors"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jordic/goics"
	m "trainingCalendar/model"
)

type Service struct {
	//TODO: dependency injection for data
}

type ScheduleService interface {
	Reschedule(f *os.File, raceDate time.Time) error
	CreateIcal(f *os.File) *os.File
	LoadCalendar(f *os.File) m.Schedule
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Reschedule(f *os.File, raceDate time.Time) error {
	fmt.Println("Rescheduling based off new race date: " + raceDate.String())
	writer := csv.NewWriter(f)
    defer writer.Flush()

    weeks := s.LoadCalendar(f)

    //TODO

    //determine day of week of race
    dayOfWeek := raceDate.Weekday()
    var prevMonday time.Time
    if dayOfWeek == time.Saturday {
    	prevMonday = raceDate.AddDate(0, 0, -5)
    } else if dayOfWeek == time.Sunday {
    	prevMonday = raceDate.AddDate(0, 0, -6)
    } else {
    	log.Printf("Only Saturday and Sunday Races are supported... Sorry")
    	return errors.New("Race date is not on a Weekend")
    }
    fmt.Println("Monday before race - " + prevMonday.Format(m.DateLayout))

    //count back from there
	for i := len(weeks)-1; i >= 0; i-- {
		weeks[i].WeekStart = prevMonday
		prevMonday = prevMonday.AddDate(0, 0, -7)
	}
	fmt.Println("New schedule: ")
    weeks.Print()

    weeks.WriteCsv()
    return nil
}

func (s *Service) CreateIcal(csvFile *os.File) *os.File {
	weeks := s.LoadCalendar(csvFile)
	b := &bytes.Buffer{}
	enc := goics.NewICalEncode(b)
	enc.Encode(weeks)
	
	calFile := "out/training.ics"
	f, err := os.Create(calFile)
	if err != nil {
		fmt.Println("Error creating ical: " + err.Error())
		log.Printf(err.Error())
	}

	w := bufio.NewWriter(f)
	b.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing ical: " + err.Error())
		log.Printf(err.Error())
	}

	w.Flush()
	if err != nil {
		fmt.Println("Error flushing writer: " + err.Error())
		log.Printf(err.Error())
	}

	return f
}

func (s *Service) LoadCalendar(f *os.File) m.Schedule {
	// Read File into a Variable
	f.Seek(0,0)
	r := csv.NewReader(f)
    if _, err := r.Read(); err != nil { //read header
    	fmt.Println("Error reading csv header: " + err.Error())
        log.Fatal(err)
    }
    lines, err := r.ReadAll()
    if err != nil {
    	fmt.Println("Error Reading Csv: " + err.Error())
        panic(err)
    }

    var sched m.Schedule

    // Loop through lines & turn into object
    for _, line := range lines {
    	var days [7]m.Event
		monday, err := time.Parse(m.DateLayout, line[0])
		if err != nil {
			monday, _ = time.Parse(m.BackupDateLayout, line[0])
		}
		for i, _ := range line[1:] {
    		days[i] = m.Event{monday.AddDate(0, 0, i), line[i+1]}
       	}
        wk := &m.Week{
            WeekStart: monday,
            Days: days,
        }
        sched = append(sched, wk)
    }

    return sched
}
