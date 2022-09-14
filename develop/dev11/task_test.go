package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wb_l2/develop/dev11/internal/calendar"
)

func TestAddDelete(t *testing.T) {
	c := calendar.NewCalendar()
	e := calendar.NewEvent()
	c.AddEventToCalendar(e)
	ok := c.DeleteEvent(e.ID)
	assert.Equal(t, true, ok)
}
