package query

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// New returns a query with substitutions replaced
func New(queryTemplate string, period string) (string, error) {
	var duration int64

	switch period {
	case "current_calendar_month":
		duration = durationForCurrentCalendarMonth()
	default:
		return "", errors.New("Invalid period")
	}

	durationString := fmt.Sprintf("%ds", duration)
	query := strings.Replace(queryTemplate, "$DURATION", durationString, -1)
	return query, nil
}

func durationForCurrentCalendarMonth() int64 {
	currentTime := time.Now()
	y := currentTime.Year()
	m := currentTime.Month()
	currentLocation := currentTime.Location()

	firstOfMonth := time.Date(y, m, 1, 0, 0, 0, 0, currentLocation)

	diff := currentTime.Sub(firstOfMonth)
	s := int64(diff.Seconds())
	return s
}
