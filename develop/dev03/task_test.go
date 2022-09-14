package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortLines(t *testing.T) {

	tests := []struct {
		name         string
		file         string
		expectedFile string
		arguments    arguments
	}{
		{
			"sort as string, no column",
			"./testdata/test_01.txt",
			"./testdata/test_01_expected.txt",
			arguments{
				sortByCol:        -1,
				sortAsNums:       false,
				sortReversed:     false,
				removeDuplicates: false,
			},
		},
		{
			"sort as string by column",
			"./testdata/test_02.txt",
			"./testdata/test_02_expected.txt",
			arguments{
				sortByCol:        1,
				sortAsNums:       false,
				sortReversed:     false,
				removeDuplicates: false,
			},
		},
		{
			"sort as number, no column",
			"./testdata/test_03.txt",
			"./testdata/test_03_expected.txt",
			arguments{
				sortByCol:        -1,
				sortAsNums:       true,
				sortReversed:     false,
				removeDuplicates: false,
			},
		},
		{
			"sort as number, by column",
			"./testdata/test_04.txt",
			"./testdata/test_04_expected.txt",
			arguments{
				sortByCol:        1,
				sortAsNums:       true,
				sortReversed:     false,
				removeDuplicates: false,
			},
		},
		{
			"sort as number, by column, reverse order, remove duplicates",
			"./testdata/test_05.txt",
			"./testdata/test_05_expected.txt",
			arguments{
				sortByCol:        1,
				sortAsNums:       true,
				sortReversed:     true,
				removeDuplicates: true,
			},
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		linesToSort, err := readLinesFromFile(tt.file)
		assert.NoError(t, err)
		expectedLines, err := readLinesFromFile(tt.expectedFile)
		assert.NoError(t, err)

		result, err := SortLines(linesToSort, tt.arguments)
		assert.NoError(t, err)

		assert.Equal(t, expectedLines, result)
	}
}
