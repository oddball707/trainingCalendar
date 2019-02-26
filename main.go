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
const backupDateLayout = "1/2/06"
const icalLayout = "20060102"
var header = []string{"StartDate", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var file = "calendar.csv"



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
