package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	tests := []struct {
		name         string
		sigDurations []time.Duration
		expected     interface{}
	}{
		{
			"case from the sample",
			[]time.Duration{
				2 * time.Hour,
				5 * time.Minute,
				1 * time.Second,
				1 * time.Hour,
				1 * time.Minute,
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		var sigs []<-chan interface{}
		for _, d := range tt.sigDurations {
			sigs = append(sigs, sig(d))
		}
		start := time.Now()
		<-or(sigs...)
		assert.LessOrEqual(t, time.Nanosecond, time.Since(start))
	}
}
