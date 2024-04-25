// Package dates provides functionality to get date ranges
// for 'thisweek', 'thismonth', and 'lastmonth' periods.
package dates

import (
	"time"
)

// DateRange holds a start (From) and an end date (To) in the format YYYY-MM-DD.
type DateRange struct {
	From string
	To   string
}

// GetDateRange calculates the date range based on the given period.
func GetDateRange(period string) DateRange {
	now := time.Now()

	switch period {
	case "month":
		return thisMonth(now)
	case "lastmonth":
		return lastMonth(now)
	case "week":
		return thisWeek(now)
	default:
		// Default to current day range
		return DateRange{From: now.Format("2006-01-02"), To: now.Format("2006-01-02")}
	}
}

func thisMonth(now time.Time) DateRange {
	year, month, _ := now.Date()
	location := now.Location()

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, location)
	lastDay := firstDay.AddDate(0, 1, -1)

	return DateRange{
		From: firstDay.Format("2006-01-02"),
		To:   lastDay.Format("2006-01-02"),
	}
}

func lastMonth(now time.Time) DateRange {
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthLastDay := lastMonthFirstDay.AddDate(0, 1, -1)

	return DateRange{
		From: lastMonthFirstDay.Format("2006-01-02"),
		To:   lastMonthLastDay.Format("2006-01-02"),
	}
}

func thisWeek(now time.Time) DateRange {
	// Start from Monday (modify according to your locale)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := now.AddDate(0, 0, offset)
	weekEnd := weekStart.AddDate(0, 0, 6)

	return DateRange{
		From: weekStart.Format("2006-01-02"),
		To:   weekEnd.Format("2006-01-02"),
	}
}
