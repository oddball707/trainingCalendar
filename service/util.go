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

func NextMonday(now ...time.Time) time.Time {
	if len(now) == 0 {
    	now[0] = time.Now()
  	}
	if now[0].Weekday() == time.Sunday {
		return now[0].AddDate(0, 0, 1)
	}
	return now[0].AddDate(0, 0, (1 + 7 - int(now[0].Weekday())%7))
}