package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name         string
		pattern      string
		files        []string
		expectedFile string
		args         arguments
	}{
		{
			"single file, default case, prefix line numbers, ignore case",
			"addsadasd",
			[]string{"./testdata/test_01.txt"},
			"./testdata/test_01_expected.txt",
			arguments{
				after:      0,
				before:     0,
				context:    0,
				count:      false,
				ignoreCase: true,
				invert:     false,
				fixed:      false,
				lineNum:    true,
			},
		},
		{
			"multiple files, only print number of matched lines, ignore case",
			"addsadasd",
			[]string{"./testdata/test_01.txt", "./testdata/test_02.txt"},
			"./testdata/test_02_expected.txt",
			arguments{
				after:      0,
				before:     0,
				context:    0,
				count:      true,
				ignoreCase: true,
				invert:     false,
				fixed:      false,
				lineNum:    false,
			},
		},
		{
			"exact match, print context",
			"0",
			[]string{"./testdata/test_03.txt"},
			"./testdata/test_03_expected.txt",
			arguments{
				after:      0,
				before:     0,
				context:    1,
				count:      false,
				ignoreCase: false,
				invert:     false,
				fixed:      true,
				lineNum:    false,
			},
		},
		{
			"invert, print 1 before and 1 after",
			"0",
			[]string{"./testdata/test_04.txt"},
			"./testdata/test_04_expected.txt",
			arguments{
				after:      1,
				before:     1,
				context:    0,
				count:      false,
				ignoreCase: false,
				invert:     true,
				fixed:      false,
				lineNum:    false,
			},
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		result, err := Grep(tt.pattern, tt.files, tt.args)
		assert.NoError(t, err)
		expectedLines := readLinesFromFile(tt.expectedFile)

		assert.Equal(t, strings.Join(expectedLines, "\n"), result)
	}
}
