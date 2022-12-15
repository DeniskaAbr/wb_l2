package utils

import (
	"dev11/api/domain"
	"time"
)

func StartEndFromTimeInterval(interval domain.TimeInterval, t time.Time) (start, end time.Time, err error) {
	currentTimestamp := t.UTC()
	currentYear, currentMonth, currentDay := currentTimestamp.Date()
	currentLocation := currentTimestamp.Location()
	currentWeekDay := int(currentTimestamp.Weekday())

	switch interval {
	case domain.Day:
		start = time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
		end = time.Date(currentYear, currentMonth, currentDay, 23, 59, 59, 999999999, currentLocation)
	case domain.Week:
		// sunday is first +1
		start = time.Date(currentYear, currentMonth, currentDay-currentWeekDay+1, 0, 0, 0, 0, currentLocation)
		end = time.Date(currentYear, currentMonth+1, currentDay+6-currentWeekDay+1, 23, 59, 59, 999999999, currentLocation)
	case domain.Month:
		start = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		end = time.Date(currentYear, currentMonth+1, 0, 23, 59, 59, 999999999, currentLocation)
	case domain.Year:
		start = time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
		end = time.Date(currentYear, 13, 0, 23, 59, 59, 999999999, currentLocation)
	}

	return start, end, nil
}
