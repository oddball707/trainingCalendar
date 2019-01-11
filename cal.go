package main

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

const dateLayout = "01/02/2006"
const icalLayout = "20060102"
var header = []string{"StartDate", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var file = "calendar.csv"

// Event is a date entry with the date and description
type Event struct {
	date   		time.Time `json:"date"`
	description string    `json:"description"`
}

//weekStart is the date of Monday
//days is a array of events, starting with monday
type Week struct {
	weekStart	time.Time   `json:"weekStart"`
	days		[7]Event	`json:"days"`
}

type Schedule []*Week

func checkError(message string, err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf(err.Error())
	}
}

func (s Schedule) Print() {
	fmt.Println("Week |  StartDate | Monday|  Tuesday  |Wednesday| Thursday | Friday | Saturday | Sunday")
	for i, week := range s {
		fmt.Print(i)
		fmt.Print("    | ")
		fmt.Print(week.weekStart.Format(dateLayout))
		fmt.Print(" | ")
		for _, day := range week.days {
			fmt.Print(day.description)
			fmt.Print("  | ")
		}
		fmt.Println(" ")
	}
}

func (s Schedule) WriteCsv() {
	file, err := os.OpenFile("calendar.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	checkError("Failed to open calendar.csv", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
    writer.Write(header)
	for _, week := range s {
		line := make([]string, 8)
		line[0] = week.weekStart.Format(dateLayout)
		for i, day := range week.days {
			line[i+1] = day.description
		}
		fmt.Println(line)
		err := writer.Write(line)
		checkError("Failed to write line", err)
	}
}

// EmitICal implements the interface for goics
func (s Schedule) EmitICal() goics.Componenter {
	component := goics.NewComponent()
	component.SetType("VCALENDAR")
	component.AddProperty("CALSCAL", "GREGORIAN")

	for _, week := range s {
		for _, day := range week.days {
			c := goics.NewComponent()
			c.SetType("VEVENT")
			k, v := goics.FormatDateField("DTEND", day.date)
			c.AddProperty(k, v)
			k, v = goics.FormatDateField("DTSTART", day.date)
			c.AddProperty(k, v)
			c.AddProperty("SUMMARY", day.description)

			component.AddComponent(c)
		}
	}
	return component
}

func reschedule(f *os.File, raceDate time.Time) {
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

func createIcal(f *os.File) {
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

func loadCalendar(f *os.File) Schedule {
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
		monday, _ := time.Parse(dateLayout, line[0])
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

func main() {

	recalcDate := flag.String("date", "", "Recalculate schedule based on date of race passed in (mm/dd/yyyy)")
	create := flag.Bool("create", false, "Create a new ical")
	filestr := flag.String("file", "", "csv file to use")
	help := flag.Bool("h", false, "Prints this help info")
	flag.Parse()

	if(*help) {
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if(*filestr != "") {
		file = *filestr
	}

	// Open CSV file
    f, err := os.Open(file)
    if err != nil {
    	fmt.Println("Failed to open csv file: " + file)
        panic(err)
    }
    defer f.Close()

	if( *recalcDate != "") {
		raceDate, err := time.Parse(dateLayout, *recalcDate)
		if err != nil {
			fmt.Println("Improperly formated date: " + *recalcDate)
			panic(err)
		}
		reschedule(f, raceDate)
	}
    if(*create) {
		createIcal(f);
	}
    if !*create && *recalcDate == "" {
    	flag.PrintDefaults()
	}
}
