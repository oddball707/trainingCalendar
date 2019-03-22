package model

import (
	"time"
)

// Event is a date entry with the date and description
type Event struct {
	Date   		time.Time `json:"date"`
	Description string    `json:"description"`
}