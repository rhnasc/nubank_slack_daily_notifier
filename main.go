package main

import (
	"os"
	"time"

	"github.com/rhnasc/nubank_api_exporter/nubank"
)

func main() {
	login := os.Getenv("NUBANK_LOGIN")
	password := os.Getenv("NUBANK_PASSWORD")

	client := nubank.NewNubankHttpClient(login, password)

	must(client.Discover())
	must(client.Authenticate())

	events, err := client.Events()
	must(err)

	filteredEvents := FilterEventsByTimeRange(
		events,
		time.Now().Add(-24*time.Hour),
		time.Now(),
	)

	webhookAddr := os.Getenv("SLACK_WEBHOOK_ADDRESS")

	slackClient := NewSlackClient(webhookAddr)
	must(slackClient.SendEvents(filteredEvents))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
