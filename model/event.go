package model

// Event is a date entry with the date and description
type Event struct {
	date   		time.Time `json:"date"`
	description string    `json:"description"`
}