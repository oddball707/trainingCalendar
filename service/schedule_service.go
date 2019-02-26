package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jordic/goics"
)

type Service struct {
	//TODO: dependency injection for data
}

type ScheduleService interface {
	Reschedule(f *os.File, raceDate time.Time)
	CreateIcal(f *os.File)
	LoadCalendar(f *os.File) Schedule
}

func newService() *Service {
	return &Service{}
}

func (s *Service) Reschedule(f *os.File, raceDate time.Time) {
	fmt.Println("Rescheduling based off new race date: " + raceDate.String())
	writer := csv.NewWriter(f)
    defer writer.Flush()

    weeks := loadCalendar(f)

    //TODO

    //determine day of week of race
    dayOfWeek := raceDate.Weekday()
    var prevMonday time.Time
    if dayOfWeek == time.Saturday {
    	prevMonday = raceDate.AddDate(0, 0, -5)
    } else if dayOfWeek == time.Sunday {
    	prevMonday = raceDate.AddDate(0, 0, -6)
    } else {
    	log.Fatalf("Only Saturday and Sunday Races are supported... Sorry")
    }
    fmt.Println("Monday before race - " + prevMonday.Format(dateLayout))

    //count back from there
	for i := len(weeks)-1; i >= 0; i-- {
		weeks[i].weekStart = prevMonday	
		prevMonday = prevMonday.AddDate(0, 0, -7)
	}
	fmt.Println("New schedule: ")
    weeks.Print()

    weeks.WriteCsv()
}

func (s *Service) CreateIcal(f *os.File) {
	weeks := loadCalendar(f)

	b := &bytes.Buffer{}
	enc := goics.NewICalEncode(b)
	enc.Encode(weeks)

	f, err := os.Create("./training.ics")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf(err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	b.WriteTo(w)
	w.Flush()
}

func (s *Service) LoadCalendar(f *os.File) Schedule {
	// Read File into a Variable
	r := csv.NewReader(f)
    if _, err := r.Read(); err != nil { //read header
    	fmt.Println(err.Error())
        log.Fatal(err)
    }
    lines, err := r.ReadAll()
    if err != nil {
    	fmt.Println(err.Error())
        panic(err)
    }

    var sched Schedule

    // Loop through lines & turn into object
    for _, line := range lines {
    	var days [7]Event
		monday, err := time.Parse(dateLayout, line[0])
		if err != nil {
			monday, _ = time.Parse(backupDateLayout, line[0])
		}
		for i, _ := range line[1:] {
    		days[i] = Event{monday.AddDate(0, 0, i), line[i+1]}
       	}
        wk := Week{
            weekStart: monday,
            days: days,
        }
        sched = append(sched, &wk)
    }

    return sched
}
