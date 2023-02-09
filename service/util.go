package service

import (
	"time"
)

func PrevMonday(day time.Time) time.Time {
	if day.Weekday() == time.Sunday {
		return day.AddDate(0, 0, -6)
	}
	return day.AddDate(0, 0, (int(day.Weekday())-1)*-1)
}

func NextMonday(now time.Time) time.Time {
	if now.Weekday() == time.Sunday {
		return now.AddDate(0, 0, 1)
	}
	return now.AddDate(0, 0, (1 + 7 - int(now.Weekday())%7))
}