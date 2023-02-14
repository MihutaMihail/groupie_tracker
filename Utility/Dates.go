package utility

import (
	"fmt"
	"log"
	"sort"
	"time"
)

const layout = "02-01-2006"

func SortDates(datesWithAsterix []string) []string {
	var finalList []string
	var dates []string
	var times []time.Time

	fmt.Println(datesWithAsterix)
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

	fmt.Print(finalList)
	return finalList
}
