package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNTPTimeWrongHost(t *testing.T) {
	_, err := GetNtpTime("wronghost")
	assert.Error(t, err)
}

func TestGetNtpTimeCorrectTime(t *testing.T) {
	ntpTime, err := GetNtpTime(host)
	assert.NoError(t, err)

	curTime := time.Now()

	assert.Equal(t, curTime.Year(), ntpTime.Year())
	assert.Equal(t, curTime.Month(), ntpTime.Month())
	assert.Equal(t, curTime.Day(), ntpTime.Day())
	assert.Equal(t, curTime.Hour(), ntpTime.Hour())
	assert.Equal(t, curTime.Minute(), ntpTime.Minute())
	assert.Equal(t, curTime.Second(), ntpTime.Second())
}
