package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackNoErrors(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"general example", "a4bc2d5e", "aaaabccddddde"},
		{"require no unpacking", "abcd", "abcd"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		result, err := Unpack(tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.output, result)
	}
}

func TestUnpackWithErrors(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"incorrect input", "45", ""},
		{"input with zero", "abc0d", ""},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		result, err := Unpack(tt.input)
		assert.Error(t, err)
		assert.Equal(t, tt.output, result)
	}
}
