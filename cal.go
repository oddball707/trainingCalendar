package main

import (
	"encoding/json"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

)

const dateLayout = "01/02/2006"
const file = "calendar.csv"

// Event is a time entry
type Event struct {
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	Description string    `json:"description"`
}

type Week struct {
	weekStart	time.Time   `json:"weekStart"`
	monday		string 		`json:"monday"`
	tuesday		string 		`json:"tuesday"`
	wednesday	string 		`json:"wednesday"`
	thursday	string 		`json:"thursday"`
	friday		string 		`json:"friday"`
	saturday	string 		`json:"saturday"`
	sunday		string 		`json:"sunday"`
}

// EmitICal implements the interface for goics
func (e Event) EmitICal(c goics.Componenter) {
	c.AddProperty("CALSCAL", "GREGORIAN")
	s := goics.NewComponent()
	s.SetType("VEVENT")
	k, v := goics.FormatDateTimeField("DTEND", e.DateEnd)
	s.AddProperty(k, v)
	k, v = goics.FormatDateTimeField("DTSTART", e.DateStart)
	s.AddProperty(k, v)
	s.AddProperty("SUMMARY", e.Description)

	c.AddComponent(s)
}


func reschedule(f *os.File, raceDate time.Time) {
	writer := csv.NewWriter(file)
    defer writer.Flush()

    weeks := loadCalendar(f)

    

}

func create(f *os.File) {


}

func loadCalendar(f *os.File) []Week {
	// Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

    var sched []Week

    // Loop through lines & turn into object
    for _, line := range lines {
        wk := Week{
            weekStart: line[0],
            monday: line[1],
         	tuesday: line[2],	
			wednesday: line[3],
			thursday: line[4],
			friday: line[5],
			saturday: line[6],
			sunday: line[7]
        }
        sched = append(sched, wk)
    }

    return sched
}

func main() {

	recalcDate := flag.String("date", "", "Recalculate schedule based on date of race passed in (mm/dd/yyyy)")
	create := flag.Bool("create", false, "Create a new ical")
	filestr := flag.String("file", "", "csv file to use")
	flag.Parse()


	if(*filestr != "") {
		file = *filestr
	}

		// Open CSV file
    f, err := os.Open(file)
    if err != nil {
        panic(err)
    }
    defer f.Close()

	if( *recalcDate == "") {
		createIcal(srv);
	} else if(*create) {
		raceDate, err := time.Parse(dateLayout, *recalcDate)
		if err != nil {
			log.Fatalf("Improperly formated date: " + *recalcDate)
		}
		reschedule(srv, raceDate)
	} else {
		flag.PrintDefaults()
	}
}
