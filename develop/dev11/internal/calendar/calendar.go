package calendar

import (
	"fmt"
	"sync"
	"time"
)

type Calendar struct {
	*sync.RWMutex
	data map[string]*Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		&sync.RWMutex{},
		make(map[string]*Event),
	}
}

func (c *Calendar) AddEventToCalendar(e *Event) {
	c.Lock()
	defer c.Unlock()

	c.data[e.ID] = e
}

func (c *Calendar) DeleteEvent(id string) (ok bool) {
	c.Lock()
	defer c.Unlock()

	if _, ok = c.data[id]; ok {
		delete(c.data, id)
	}
	return ok
}

func (c *Calendar) getEventById(id string) *Event {
	c.RLock()
	defer c.RUnlock()

	return c.data[id]
}

func (c *Calendar) UpdateEvent(e *Event) error {
	event := c.getEventById(e.ID)
	if event == nil {
		return fmt.Errorf("event with ID=%s not found", e.ID)
	}

	event.Lock()
	defer event.Unlock()

	event.Title = e.Title
	event.Date = e.Date
	event.Description = e.Description
	event.UserID = e.UserID

	return nil
}

func (c *Calendar) GetEventsInTimeRange(userID string, start, end time.Time) (e []*Event) {
	c.RLock()
	defer c.RUnlock()

	for _, v := range c.data {
		if v.UserID == userID && ((v.Date == start || v.Date.After(start)) && v.Date.Before(end)) {
			e = append(e, v)
		}
	}

	return e
}
