package model

//weekStart is the date of Monday
//days is a array of events, starting with monday
type Week struct {
	weekStart	time.Time   `json:"weekStart"`
	days		[7]Event	`json:"days"`
}