package main

import (
	"log"
	"net/http"

	h "trainingCalendar/handler"
	s "trainingCalendar/service"
)

func main() {
	// initialize service layer
	srv := s.NewService()

	hnd := h.NewHandler(srv)


	// create http server
	http.HandleFunc("/health", hnd.HealthHandler)
	http.HandleFunc("/readiness", hnd.ReadinessHandler)
	http.HandleFunc("/create", hnd.CreateSchedule)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	log.Printf("Service started on 0.0.0.0:8080")

}
