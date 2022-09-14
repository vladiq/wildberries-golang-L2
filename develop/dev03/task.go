package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type arguments struct {
	sortByCol        int
	sortAsNums       bool
	sortReversed     bool
	removeDuplicates bool
}

type keyValue struct {
	key   string
	value interface{}
}

func main() {
	sortByCol := flag.Int("k", -1, "columns to sort by. Default: -1")
	sortAsNums := flag.Bool("n", false, "whether to sort as numbers. Default: false")
	sortReversed := flag.Bool("r", false, "sort in reverse order. Default: false")
	removeDuplicates := flag.Bool("u", false, "remove duplicate lines. Default: false")
	flag.Parse()

	args := arguments{*sortByCol, *sortAsNums, *sortReversed, *removeDuplicates}

	fileName := "/home/vlad/GolandProjects/wb_l2/develop/dev03/testdata/test_01.txt"
	lines, err := readLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}

	sortedLines, err := SortLines(lines, args)
	if err != nil {
		panic(err)
	}
	fmt.Println(sortedLines)
}

func SortLines(lines []string, args arguments) ([]string, error) {
	if args.removeDuplicates {
		lines = removeDuplicateStr(lines)
	}

	if args.sortByCol == -1 {
		return sortByNoColumn(lines, args)
	}

	return sortByColumn(lines, args)
}

func sortByColumn(lines []string, args arguments) ([]string, error) {
	sliceForSorting := make([]keyValue, 0, len(lines))

	// split lines
	for _, line := range lines {
		split := strings.Split(line, " ")
		if args.sortAsNums {
			if val, err := strconv.Atoi(split[args.sortByCol]); err != nil {
				return nil, errors.New("cannot sort a column as a number, because it contains letters")
			} else {
				sliceForSorting = append(sliceForSorting, keyValue{line, val})
			}
		} else {
			sliceForSorting = append(sliceForSorting, keyValue{line, split[args.sortByCol]})
		}
	}

	if args.sortReversed {
		sort.Slice(sliceForSorting, func(i, j int) bool {
			if args.sortAsNums {
				return sliceForSorting[i].value.(int) > sliceForSorting[j].value.(int)
			} else {
				return sliceForSorting[i].value.(string) > sliceForSorting[j].value.(string)
			}
		})
	} else {
		sort.Slice(sliceForSorting, func(i, j int) bool {
			if args.sortAsNums {
				return sliceForSorting[i].value.(int) < sliceForSorting[j].value.(int)
			} else {
				return sliceForSorting[i].value.(string) < sliceForSorting[j].value.(string)
			}
		})
	}

	result := make([]string, 0, len(lines))
	for _, s := range sliceForSorting {
		result = append(result, s.key)
	}
	return result, nil
}

func sortByNoColumn(lines []string, args arguments) ([]string, error) {
	switch args.sortAsNums {
	case true:
		// convert strings to nums
		ints := make([]int, 0, len(lines))
		for _, line := range lines {
			lineAsInt, err := strconv.Atoi(line)
			if err != nil {
				return nil, errors.New("cannot sort string with letters as a number")
			}
			ints = append(ints, lineAsInt)
		}

		if args.sortReversed {
			sort.Sort(sort.Reverse(sort.IntSlice(ints)))
		} else {
			sort.Ints(ints)
		}

		// convert nums back to strings
		result := make([]string, 0, len(ints))
		for _, line := range ints {
			lineAsStr := strconv.Itoa(line)
			result = append(result, lineAsStr)
		}
		return result, nil

	default:
		if args.sortReversed {
			sort.Sort(sort.Reverse(sort.StringSlice(lines)))
		} else {
			sort.Strings(lines)
		}
		return lines, nil
	}
}

func removeDuplicateStr(lines []string) []string {
	var result []string
	uniqueLines := make(map[string]struct{})

	for _, item := range lines {
		if _, ok := uniqueLines[item]; !ok {
			uniqueLines[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func readLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
