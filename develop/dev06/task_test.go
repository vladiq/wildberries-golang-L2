package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	tests := []struct {
		name     string
		args     arguments
		input    []string
		expected []string
	}{
		{
			"default delimiter, first field",
			arguments{
				fields:    "1",
				delimiter: "\t",
				separated: false,
			},
			[]string{
				"123\t321",
				"abs\tbab",
				"456",
			},
			[]string{
				"123",
				"abs",
				"456",
			},
		},

		{
			"custom delimiter, multiple fields",
			arguments{
				fields:    "1,3-5",
				delimiter: " ",
				separated: false,
			},
			[]string{
				"1 2 3 4 5 6",
				"1 2 3 4 5",
				"1 2 3 4",
				"1 2 3",
				"1 2",
				"1",
			},
			[]string{
				"1 3 4 5",
				"1 3 4 5",
				"1 3 4",
				"1 3",
				"1",
				"1",
			},
		},

		{
			"only lines that contain delimiter",
			arguments{
				fields:    "1,3-5",
				delimiter: " ",
				separated: true,
			},
			[]string{
				"1 2 3 4 5 6",
				"12345",
				"1 2 3 4",
				"123",
				"1 2",
				"1",
			},
			[]string{
				"1 3 4 5",
				"1 3 4",
				"1",
			},
		},
	}
	for _, tt := range tests {
		t.Log(tt.name)
		input := []byte(strings.Join(tt.input, "\n"))

		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		_, err = w.Write(input)
		if err != nil {
			t.Error(err)
		}
		w.Close()

		stdin := os.Stdin
		os.Stdin = r

		actualOutput, err := Cut(tt.args)
		assert.NoError(t, err)
		assert.Equal(t, strings.Join(tt.expected, "\n")+"\n", actualOutput)
		os.Stdin = stdin
	}

}
