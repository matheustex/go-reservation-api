package util

import (
	"time"
)

type DateInterval struct {
	StartDate string
	EndDate   string
}

const (
	// ISO8601 format
	ISO8601 = "2006-01-02T15:04:05.999Z"
)

func Now() string {
	return time.Now().UTC().Format(ISO8601)
}

func IsInvalidDate(start string, end string) bool {
	startDate, _ := time.Parse(ISO8601, start)
	endDate, _ := time.Parse(ISO8601, end)

	return startDate.Before(time.Now().UTC()) || endDate.Before(time.Now().UTC())
}

func IsDateBetween(interval DateInterval, date string) bool {
	startDate, _ := time.Parse(ISO8601, interval.StartDate)
	endDate, _ := time.Parse(ISO8601, interval.EndDate)
	Date, _ := time.Parse(ISO8601, date)

	return Date.After(startDate) && Date.Before(endDate)
}

func IsIntervalBetween(firstInterval, secondInterval DateInterval) bool {
	return IsDateBetween(firstInterval, secondInterval.StartDate) &&
		IsDateBetween(firstInterval, secondInterval.EndDate)
}

func IsDateBooked(reservationInterval, refInterval DateInterval) bool {
	return IsDateBetween(reservationInterval, refInterval.StartDate) ||
		IsDateBetween(reservationInterval, refInterval.EndDate) ||
		IsIntervalBetween(reservationInterval, refInterval)
}
