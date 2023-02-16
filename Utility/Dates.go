package utility

import (
	"log"
	"sort"
	"time"
)

const layout = "02-01-2006"

func SortDates(datesWithAsterix []string) []string {
	var finalList []string
	var dates []string
	var times []time.Time

	// enlève les asterix
	for _, date := range datesWithAsterix {
		if date[0] == '*' {
			dates = append(dates, date[1:])
		} else {
			dates = append(dates, date)
		}
	}

	// transforme en dates
	for _, date := range dates {
		tm, err := time.Parse(layout, date)
		if err != nil {
			log.Println(err)
		}
		times = append(times, tm)
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	for _, t := range times {
		tm := t.Format(layout)
		finalList = append(finalList, tm)
	}

	return finalList
}
