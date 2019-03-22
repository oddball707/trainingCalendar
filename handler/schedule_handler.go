package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Fprintln(w, "Healthy")
}

func(h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ready")
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
		http.Error(w, err.Error(), 500)
		return
	}

	raceDate, err := time.Parse(m.DateLayout, msg.Date)
	if err != nil {
		fmt.Println("Improperly formated date: " + msg.Date)
		panic(err)
	}

	// Open CSV file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Failed to open csv file: " + file)
		panic(err)
	}
	defer f.Close()

	h.service.Reschedule(f, raceDate)
	calFile := h.service.CreateIcal(f)
	defer calFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=" + out)
	w.Header().Set("Content-Type", "text/calendar")
	//stream the body to the client without fully loading it into memory
	http.ServeFile(w, r, calFile.Name())
}

