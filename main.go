package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	h "trainingCalendar/handler"
	m "trainingCalendar/model"
	s "trainingCalendar/service"
)

func main() {
	cmd := flag.Bool("cmd", false, "Use this to run locally")
	recalcDate := flag.String("date", "", "Recalculate schedule based on date of race passed in (mm/dd/yyyy)")
	help := flag.Bool("h", false, "Prints this help info")
	flag.Parse()

	if(*help) {
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// initialize service layer
	srv := s.NewService()
	hnd := h.NewHandler(srv)

	if *cmd {
		if( *recalcDate != "") {
			raceDate, err := time.Parse(m.DateLayout, *recalcDate)
			if err != nil {
				fmt.Println("Improperly formated date: " + *recalcDate)
				panic(err)
			}

			race := &m.Race{
				RaceDate: raceDate,
				RaceType: m.Dynamic,
			}

			options := &m.Options{
				WeeklyMileage: 50,
				RestDays: 2,
				BackToBacks: true,
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
		serve(srv, hnd)
	}
}

func serve(srv *s.Service, hnd *h.Handler) {
	router := hnd.NewRouter()
//	router.Get("/", http.FileServer(http.Dir("./web")))
	hnd.FileServer(router, "/", http.Dir("./web"))

	http.Handle("/", router)
	port := getEnv("PORT", "8080")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}
	log.Printf("Service started on 0.0.0.0:8080")

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
