package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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
		if *recalcDate == "" {
			fmt.Println("When is your race (Or how long should this schedule continue)? (enter in format mm/dd/yyyy)")
			fmt.Scanln(recalcDate)
		}
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

		var weeklyMileage, restDays, wowIncrease, restWeekFreq, restWeekLevel int
		b2b := false
		var b2bString string

		if rType == m.Dynamic {
			fmt.Println("How many miles are you running per week now?")
			fmt.Scanln(&weeklyMileage)

			if weeklyMileage < 15 {
				fmt.Println("You can probably start at 15...")
				weeklyMileage = 15
			}

			fmt.Println("How rest days do you prefer each week? (1-4)")
			fmt.Scanln(&restDays)

			if restDays < 1 || restDays > 4 {
				fmt.Println("Not recommended, you get 2 rest days")
				restDays = 2
			}

			fmt.Println("Do you want back to back long runs on the weekend? (y/N)")
			fmt.Scanln(&b2bString)

			if strings.ToUpper(b2bString) == "Y" {
				b2b = true
			}

			fmt.Println("How often should a rest week be scheduled? (Default every 3 weeks, 2-6 required)")
			fmt.Scanln(&restWeekFreq)

			if restWeekFreq < 2 || restWeekFreq > 6 {
				fmt.Println("Not recommended, revert to default")
				restWeekFreq = 3
			}

			fmt.Println("How much mileage increase do you want week over week? (Default 10%, 5-25% required)")
			fmt.Scanln(&wowIncrease)

			if wowIncrease < 5 || wowIncrease > 25 {
				fmt.Println("Not recommended, revert to default")
				wowIncrease = 10
			}

			fmt.Println("What percent of a standard week do you want a rest week to be? (Default 70%, 10-90% required)")
			fmt.Scanln(&restWeekLevel)

			if restWeekLevel < 10 || wowIncrease > 90 {
				fmt.Println("Not recommended, revert to default")
				restWeekLevel = 70
			}
		}

		options := &m.Options{
			WeeklyMileage: weeklyMileage,
			RestDays:      restDays,
			BackToBacks:   b2b,
			WowIncrease:   wowIncrease,
			RestWeekFreq:  restWeekFreq,
			RestWeekLevel: restWeekLevel,
		}

		fmt.Printf("Creating calendar with your options: %v\n", options)

		var filePath string
		fmt.Println("Where do you want to save your calendar? (Default ./)")
		fmt.Scanln(&filePath)

		if _, err := os.Stat(filePath); err != nil {
			fmt.Println("Path does not exist, default to .")
			filePath = "."
		}

		calFile, err := srv.CreateIcal(race, options, filePath)
		if err != nil {
			log.Printf("Issue creating ical: %v", err)
		}
		defer calFile.Close()
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
