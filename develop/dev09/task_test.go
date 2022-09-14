package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadWebsite(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"picture of hot dog",
			"https://www.pictureofhotdog.com/about",
			3,
		},
		{
			"japanese photographer",
			"http://www.hisadomi.com/",
			4,
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		n, err := DownloadWebsite(tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, n)
	}
}
