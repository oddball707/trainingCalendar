package main

import (
	"log"
	"net/http"
	"os"
	h "trainingCalendar/handler"
	s "trainingCalendar/service"

	"github.com/gorilla/mux"
)

func main() {
	// initialize service layer
	srv := s.NewService()
	hnd := h.NewHandler(srv)

	router := mux.NewRouter()
	router.Use(CORS)

	router.HandleFunc("/api/health", hnd.HealthHandler)
	router.HandleFunc("/api/readiness", hnd.ReadinessHandler)
	router.HandleFunc("/api/create", hnd.CreateIcal)
	router.HandleFunc("/api/show", hnd.CreateSchedule)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	http.Handle("/", router)
	port := getEnv("PORT", "8080")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}
	log.Printf("Service started on 0.0.0.0:8080")

}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
