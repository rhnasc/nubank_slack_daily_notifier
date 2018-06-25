package main

import (
	"time"

	"github.com/rhnasc/nubank_api_exporter/nubank"
)

func FilterEventsByTimeRange(events []*nubank.Event, from time.Time, to time.Time) []*nubank.Event {
	filteredEvents := []*nubank.Event{}

	for _, event := range events {
		onRange := event.Time.After(from) && event.Time.Before(to)

		if onRange {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents
}
