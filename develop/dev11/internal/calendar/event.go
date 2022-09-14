package calendar

import (
	"sync"
	"time"
)

type Event struct {
	sync.Mutex
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func NewEvent() *Event {
	return &Event{}
}
