package main

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/rhnasc/nubank_api_exporter/nubank"
)

const (
	nubankChannel  = "#nubank"
	nubankEmoji    = ":nubank:"
	nubankUsername = "Nubank"
	nubankColor    = "#612F74"
)

type slackClient struct {
	addr string
}

func NewSlackClient(addr string) *slackClient {
	return &slackClient{
		addr: addr,
	}
}

func (s *slackClient) SendEvents(events []*nubank.Event) error {

	payload := slack.Payload{}

	payload.Channel = nubankChannel

	payload.Username = nubankUsername
	payload.IconEmoji = nubankEmoji

	attachments := []slack.Attachment{}
	for _, event := range events {
		color := nubankColor
		title := event.Description
		description := fmt.Sprintf("R$ %d,%d", event.Amount/100, event.Amount%100)

		attachment := slack.Attachment{
			Color: &color,
			Title: &title,
			Text:  &description,
		}

		attachments = append(attachments, attachment)
	}

	payload.Attachments = attachments

	errs := slack.Send(s.addr, "", payload)

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
