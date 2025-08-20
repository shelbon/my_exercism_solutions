package booking

import (
	"fmt"
	"time"
)

const (
	layoutISO         = "1/2/2006 15:04:05"
	layoutRFC822      = "January 2, 2006 15:04:05"
	layoutRFC1123     = "Monday, January 2, 2006 15:04:05"
	layoutDescription = "Monday, January 2, 2006, at 15:04"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	scheduleISO, err := time.Parse(layoutISO, date)
	if err != nil {
		scheduleRFC822, err := time.Parse(layoutRFC822, date)
		if err != nil {
			scheduleRFC1123, err := time.Parse(layoutRFC1123, date)
			if err != nil {
				fmt.Println("Could not parse time:", err)
			}
			return scheduleRFC1123

		}
		return scheduleRFC822
	}
	return scheduleISO
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()
	schedule := Schedule(date)
	if schedule.Before(now) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	schedule := Schedule(date)
	hours := schedule.Hour()
	if hours >= 12 && hours < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format(layoutDescription))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
}
