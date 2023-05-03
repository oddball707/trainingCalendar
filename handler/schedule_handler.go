package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	m "github.com/oddball707/trainingCalendar/model"
	s "github.com/oddball707/trainingCalendar/service"
)

const out = "schedule.ics"

type Handler struct {
	service s.ScheduleService
}

type CreateRequest struct {
	Date     string    `json:"date"`
	RaceType string    `json:"type"`
	Options  m.Options `json:"options"`
}

func NewHandler(service s.ScheduleService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Healthy")
	http.StatusText(200)
}

func (h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Ready")
	http.StatusText(200)

}

func (h *Handler) CreateIcal(w http.ResponseWriter, r *http.Request) {

	race, options, err := parseCreateReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	calFile, err := h.service.CreateIcal(race, options)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer calFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+out)
	w.Header().Set("Content-Type", "text/calendar")
	//stream the body to the client without fully loading it into memory
	http.ServeFile(w, r, calFile.Name())
}

func (h *Handler) CreateSchedule(w http.ResponseWriter, r *http.Request) {

	race, options, err := parseCreateReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	schedule, err := h.service.GetSchedule(race, options)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}

func parseCreateReq(r *http.Request) (*m.Race, *m.Options, error) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	// Unmarshal
	var msg CreateRequest
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Print("Error Unmarshalling request - ", err)
		return nil, nil, err
	}

	raceDate, err := time.Parse(m.DateLayout, msg.Date)
	if err != nil {
		raceDate, err = time.Parse(m.BackupDateLayout, msg.Date)
		if err != nil {
			log.Print("Improperly formated date: " + msg.Date)
			return nil, nil, err
		}
	}

	raceTypeInt, err := strconv.Atoi(msg.RaceType)
	if err != nil {
		log.Print("Invalid race type, expected integer")
		return nil, nil, err
	}
	raceType := m.RaceType(raceTypeInt)
	race := &m.Race{
		RaceDate: raceDate,
		RaceType: raceType,
	}
	options := &msg.Options

	return race, options, nil
}
