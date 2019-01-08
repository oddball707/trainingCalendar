package cal

import "github.com/jordic/goics"

const dateLayout = "01/02/2006"

// Event is a time entry
type Event struct {
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	Description string    `json:"description"`
}

// EmitICal implements the interface for goics
func (e Event) EmitICal(c goics.Componenter) {
	c.AddProperty("CALSCAL", "GREGORIAN")
	s := goics.NewComponent()
	s.SetType("VEVENT")
	k, v := goics.FormatDateTimeField("DTEND", e.DateEnd)
	s.AddProperty(k, v)
	k, v = goics.FormatDateTimeField("DTSTART", e.DateStart)
	s.AddProperty(k, v)
	s.AddProperty("SUMMARY", e.Description)

	c.AddComponent(s)
}