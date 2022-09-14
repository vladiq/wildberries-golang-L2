package server

import (
	"time"
	cal "wb_l2/develop/dev11/internal/calendar"
)

type EventModel struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (e *EventModel) modelToEvent() *cal.Event {
	event := cal.NewEvent()
	event.ID = e.ID
	event.Title = e.Title
	event.UserID = e.UserID
	event.Description = e.Description
	event.Date, _ = time.Parse("2006-01-02", e.Date)
	return event
}

func (e *EventModel) validate() bool {
	if _, err := time.Parse("2006-01-02", e.Date); err != nil {
		return false
	}
	wrongInput := e.UserID == "" || e.Title == ""
	return !wrongInput
}
