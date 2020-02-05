package main

import (
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

const (
	maxResultsLen      = 250
	calendarID         = "primary"
	hasDeletedCalendar = false
	order              = "startTime"
)

func List(srv *calendar.Service) {
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarID).ShowDeleted(hasDeletedCalendar).
		SingleEvents(true).TimeMin(t).MaxResults(maxResultsLen).OrderBy(order).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	if len(events.Items) != 0 {
		for _, item := range events.Items {
			log.Printf("%v\n", item)
		}
	}
}
