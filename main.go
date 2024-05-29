package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"strings"

	h "github.com/oddball707/trainingCalendar/handler"
	m "github.com/oddball707/trainingCalendar/model"
	s "github.com/oddball707/trainingCalendar/service"
)

func main() {
	cmd := flag.Bool("cmd", false, "Use this to run locally")
	recalcDate := flag.String("date", "", "Recalculate schedule based on date of race passed in (mm/dd/yyyy)")
	help := flag.Bool("h", false, "Prints this help info")
	flag.Parse()

	if *help {
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// initialize service layer
	srv := s.NewService()

	if *cmd {
		if *recalcDate != "" {
			raceDate, err := time.Parse(m.DateLayout, *recalcDate)
			if err != nil {
				fmt.Println("Improperly formated date: " + *recalcDate)
				panic(err)
			}

			fmt.Println("What type of schedule do you want to generate?")
			fmt.Println("1 - Half Marathon")
			fmt.Println("2 - Marathon")
			fmt.Println("3 - 50k")
			fmt.Println("4 - 50M")
			fmt.Println("5 - 100k")
			fmt.Println("6 - 100M")
			fmt.Println("7 - Dynamic")
			fmt.Println("Enter number of your choice:")

			var rType m.RaceType
			fmt.Scanln(&rType)

			race := &m.Race{
				RaceDate: raceDate,
				RaceType: m.RaceType(rType),
			}

			var weeklyMileage, restDays int
			b2b := false
			var b2bString string

			if rType == m.Dynamic {
				fmt.Println("How many miles are you running per week now?")
				fmt.Scanln(&weeklyMileage)

				fmt.Println("How rest days do you prefer each week? (1-4)")
				fmt.Scanln(&restDays)

				fmt.Println("Do you want back to back long runs on the weekend? (Y/N)")
				fmt.Scanln(&b2bString)
								
				if strings.ToUpper(b2bString) == "Y" {
					b2b = true
				}
			}

			options := &m.Options{
				WeeklyMileage: weeklyMileage,
				RestDays:      restDays,
				BackToBacks:   b2b,
			}

			calFile, err := srv.CreateIcal(race, options)
			if err != nil {
				log.Printf("Issue creating ical: %v", err)
			}
			defer calFile.Close()
		} else {
			fmt.Println("date required")
			flag.PrintDefaults()
		}

	} else {
		serve(srv)
	}
}

func serve(srv *s.Service) {
	hnd := h.NewHandler(srv)
	router := hnd.NewRouter()
	hnd.FileServer(router, "/", http.Dir("./web"))

	http.Handle("/", router)
	port := getEnv("PORT", "8080")

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}
	fmt.Println("Service started on 0.0.0.0:8080")

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
