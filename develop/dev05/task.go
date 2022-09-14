package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type arguments struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	after := flag.Int("A", 0, "print N lines after match. Default: 0")
	before := flag.Int("B", 0, "print N lines before match. Default: 0")
	context := flag.Int("C", 0, "print N lines before and after match. Default: 0")
	count := flag.Bool("c", false, "print only number of matching lines for each file. Default: false")
	ignoreCase := flag.Bool("i", false, "ignore case. Default: false")
	invert := flag.Bool("v", false, "print only non-matching lines. Default: false")
	fixed := flag.Bool("F", false, "interpret patterns as fixed strings, not regular expressions. Default: false")
	lineNum := flag.Bool("n", false, "prefix each line with its line number. Default: false")
	flag.Parse()

	args := arguments{*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum}

	input := []string{"/home/vlad/GolandProjects/wb_l2/develop/dev05/testdata/test_01.txt"}
	pattern := "2034323"
	out, err := Grep(pattern, input, args)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func Grep(pattern string, filePaths []string, args arguments) (string, error) {
	filesLines, err := getLinesFromEachFile(filePaths)
	if err != nil {
		return "", err
	}

	matchesPerFile := make(map[string][]int, len(filePaths))
	for file, lines := range filesLines {
		oneFileMatches := grepOneFile(pattern, lines, args)
		matchesPerFile[file] = oneFileMatches
	}

	formattedResult, err := formatResult(filePaths, filesLines, matchesPerFile, args)
	return formattedResult, nil
}

func grepOneFile(pattern string, lines []string, args arguments) []int {
	var matchedLines []int

	for i, line := range lines {
		switch args.ignoreCase {
		case true:
			if args.fixed && strings.ToLower(line) == strings.ToLower(pattern) {
				matchedLines = append(matchedLines, i)
			} else if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				matchedLines = append(matchedLines, i)
			}
		default:
			if args.fixed && line == pattern {
				matchedLines = append(matchedLines, i)
			} else if strings.Contains(line, pattern) {
				matchedLines = append(matchedLines, i)
			}
		}
	}

	return matchedLines
}

func formatResult(filePaths []string, filesLines map[string][]string, matchesPerFile map[string][]int, args arguments) (string, error) {
	var sb strings.Builder

	if args.count {
		for _, file := range filePaths {
			output := file + ": " + strconv.Itoa(len(matchesPerFile[file])) + "\n"
			sb.Write([]byte(output))
		}
		return sb.String(), nil
	}

	for _, file := range filePaths {
		linesToOutput := make(map[int]struct{}, len(filesLines[file]))
		if args.invert {
			for i := 0; i < len(filesLines[file]); i++ {
				linesToOutput[i] = struct{}{}
			}
			for _, match := range matchesPerFile[file] {
				delete(linesToOutput, match)
			}
			matchesPerFile[file] = make([]int, 0, len(linesToOutput))
			for key, _ := range linesToOutput {
				matchesPerFile[file] = append(matchesPerFile[file], key)
			}
		} else {
			for _, match := range matchesPerFile[file] {
				linesToOutput[match] = struct{}{}
			}
		}

		for _, i := range matchesPerFile[file] {
			for ii := i - args.context; ii <= i+args.context; ii++ {
				if ii >= 0 && ii < len(filesLines[file]) {
					linesToOutput[ii] = struct{}{}
				}
			}

			for ii := i - args.before; ii <= i; ii++ {
				if ii >= 0 {
					linesToOutput[ii] = struct{}{}
				}
			}

			for ii := i; ii <= i+args.after; ii++ {
				if ii < len(filesLines[file]) {
					linesToOutput[ii] = struct{}{}
				}
			}
		}

		var outputLinesSorted []int
		for k, _ := range linesToOutput {
			outputLinesSorted = append(outputLinesSorted, k)
		}
		sort.Ints(outputLinesSorted)

		for _, i := range outputLinesSorted {
			if args.lineNum {
				sb.Write([]byte(strconv.Itoa(i) + " "))
			}
			sb.Write([]byte(filesLines[file][i] + "\n"))
		}
	}
	return sb.String(), nil
}

func getLinesFromEachFile(filePaths []string) (map[string][]string, error) {
	filesContents := make(map[string][]string, len(filePaths))
	for _, fp := range filePaths {
		filesContents[fp] = readLinesFromFile(fp)
	}
	return filesContents, nil
}

func readLinesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
