package service

import (
	"testing"

	"github.com/matheustex/go-reservation-api/util"
	"github.com/stretchr/testify/assert"
)

func TestIsInvalidDate(t *testing.T) {
	start := "2020-12-02T15:05:05.999Z"
	end := "2020-12-02T15:06:05.999Z"

	assert.True(t, util.IsInvalidDate(start, end))
}

func TestIsDateBetween(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"
	ref := "2020-12-02T15:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	assert.True(t, util.IsDateBetween(interval, ref))
}

func TestIsIntervalBetween(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	refInterval := util.DateInterval{
		StartDate: "2020-12-02T15:06:05.999Z",
		EndDate:   "2020-12-02T15:36:05.999Z",
	}

	assert.True(t, util.IsIntervalBetween(interval, refInterval))
}

func TestIsDateBooked(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	refInterval := util.DateInterval{
		StartDate: "2020-12-03T15:06:05.999Z",
		EndDate:   "2020-12-03T15:36:05.999Z",
	}

	assert.False(t, util.IsDateBooked(interval, refInterval))
}

func TestIsDateBookedWithEarlierStartDate(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	refInterval := util.DateInterval{
		StartDate: "2020-12-02T15:06:05.999Z",
		EndDate:   "2020-12-03T15:36:05.999Z",
	}

	assert.True(t, util.IsDateBooked(interval, refInterval))
}

func TestIsDateBookedWithLaterEndDate(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	refInterval := util.DateInterval{
		StartDate: "2020-12-01T15:06:05.999Z",
		EndDate:   "2020-12-02T15:36:05.999Z",
	}

	assert.True(t, util.IsDateBooked(interval, refInterval))
}

func TestIsDateBookedWithShortDates(t *testing.T) {
	start := "2020-12-02T14:05:05.999Z"
	end := "2020-12-02T16:06:05.999Z"

	interval := util.DateInterval{
		StartDate: start,
		EndDate:   end,
	}

	refInterval := util.DateInterval{
		StartDate: "2020-12-02T15:06:05.999Z",
		EndDate:   "2020-12-02T15:36:05.999Z",
	}

	assert.True(t, util.IsDateBooked(interval, refInterval))
}
