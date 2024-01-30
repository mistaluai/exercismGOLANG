package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var parseLayout string = "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(parseLayout, date)
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var layout string = "January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return time.Now().After(parsedTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var layout string = "Monday, January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	if parsedTime.Hour() >= 12 && parsedTime.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var layout string = "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	var returnString string = fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		parsedTime.Weekday().String(),
		parsedTime.Month().String(),
		parsedTime.Day(),
		parsedTime.Year(),
		parsedTime.Hour(),
		parsedTime.Minute())
	return returnString
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
