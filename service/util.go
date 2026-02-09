package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const pacePlaceholder = "<rp>"

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

func SetDescription(desc string, raceMiles float64, goalTime float64) string {
	goalPace := goalTime / raceMiles
	paceString := FloatToPace(goalPace)

	speedModifiers, desc := ParseSpeed(desc)
	if len(speedModifiers) > 0 {
		for _, speedModifier := range speedModifiers {
			modifiedPace := goalPace / (float64(speedModifier) / 100)
			paceString = FloatToPace(modifiedPace)

			desc = strings.Replace(desc, pacePlaceholder, paceString, 1)
		}
	}

	return desc
}

func FloatToPace(pace float64) string {
	minutes := int(pace)
	seconds := int((pace - float64(minutes)) * 60)
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func ParseSpeed(desc string) ([]int, string) {
	speeds := []int{}
	// find opening <
	for i, char := range desc {
		if char == '<' {
			// parse until % and return the number
			for j, char := range desc[i:] {
				if char == '%' {
					speedStr := desc[i+1 : i+j]
					speed, err := strconv.Atoi(speedStr)
					if err != nil {
						log.Print("Error parsing speed from description: " + err.Error())
						return nil, desc
					}
					// remove the percentage from the original string
					desc = desc[:i+1] + desc[i+j+1:]
					speeds = append(speeds, speed)
					break
				}
			}
		}
	}
	return speeds, desc
}
