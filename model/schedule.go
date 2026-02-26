package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jordic/goics"
)

var DateLayout = "01/02/2006"
var Header = []string{"StartDate", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var BackupDateLayout = "1/2/06"

// A Schedule is an array of Weeks
type Schedule []*Week

func (s Schedule) Print() {
	fmt.Println("Week |  StartDate | Monday|  Tuesday  |Wednesday| Thursday | Friday | Saturday | Sunday")
	for i, week := range s {
		fmt.Print(i)
		fmt.Print("    | ")
		fmt.Print(week.WeekStart.Format(DateLayout))
		fmt.Print(" | ")
		for _, day := range week.Days {
			fmt.Print(day.Title)
			fmt.Print("  | ")
		}
		fmt.Println(" ")
	}
}

func (s Schedule) WriteCsv() {
	file, err := os.OpenFile("data/calendar.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Failed to open calendar.csv: " + err.Error())
		log.Fatalf(err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(Header)
	if err != nil {
		fmt.Println("Failed to write header: " + err.Error())
		log.Fatalf(err.Error())
	}
	for _, week := range s {
		line := make([]string, 8)
		line[0] = week.WeekStart.Format(DateLayout)
		for i, day := range week.Days {
			line[i+1] = day.Description
		}
		fmt.Println(line)
		err := writer.Write(line)
		if err != nil {
			fmt.Println("Failed to write line: " + err.Error())
			log.Fatalf(err.Error())
		}
	}
}

// EmitICal implements the interface for goics
func (s Schedule) EmitICal() goics.Componenter {
	component := goics.NewComponent()
	component.SetType("VCALENDAR")
	component.AddProperty("CALSCAL", "GREGORIAN")

	for _, week := range s {
		for _, day := range week.Days {
			if day.Date.Before(time.Now()) {
				continue
			}
			c := goics.NewComponent()
			c.SetType("VEVENT")
			k, v := goics.FormatDateField("DTEND", day.Date)
			c.AddProperty(k, v)
			k, v = goics.FormatDateField("DTSTART", day.Date)
			c.AddProperty(k, v)
			c.AddProperty("SUMMARY", day.Title)
			c.AddProperty("DESCRIPTION", day.Description)
			component.AddComponent(c)
		}
	}
	return component
}
