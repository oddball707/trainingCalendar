package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	m "trainingCalendar/model"
	s "trainingCalendar/service"
)

const file = "data/calendar.csv"
const out = "schedule.ics"

type Handler struct {
	service s.ScheduleService
}

type DateRequest struct {
	Date	string	`json:"date"`
}

func NewHandler(service s.ScheduleService) *Handler {
	return &Handler{
		service: service,
	}
}


func(h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Healthy")
	http.StatusText(200)
}

func(h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Ready")
	http.StatusText(200)

}

func(h *Handler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg DateRequest
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Printf("Error Unmarshalling request", err)
		http.Error(w, err.Error(), 500)
		return
	}

	raceDate, err := time.Parse(m.DateLayout, msg.Date)
	if err != nil {
		raceDate, err = time.Parse(m.BackupDateLayout, msg.Date)
		if err != nil {
			log.Print("Improperly formated date: " + msg.Date)
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// Open CSV file
	f, err := os.Open(file)
	if err != nil {
		log.Print("Failed to open csv file: " + file)
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	err = h.service.Reschedule(f, raceDate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	calFile := h.service.CreateIcal(f)
	defer calFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=" + out)
	w.Header().Set("Content-Type", "text/calendar")
	//stream the body to the client without fully loading it into memory
	http.ServeFile(w, r, calFile.Name())
}

