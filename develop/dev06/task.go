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

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type arguments struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	fields := flag.String("f", "", "select only these fields. Default: \"\"(all fields)")
	delimiter := flag.String("d", "\t", "field delimiter. Default: TAB")
	separated := flag.Bool("s", false, "print only lines that contain specified delimiter. Default: false")

	args := arguments{*fields, *delimiter, *separated}
	if result, err := Cut(args); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}

func Cut(args arguments) (string, error) {
	if args.fields == "" {
		return "", errors.New("you must specify the field numbers")
	}

	fieldIds, err := getFieldIds(args.fields)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if args.separated && !strings.Contains(line, args.delimiter) {
			continue
		}

		inpLineSplit := strings.Split(line, args.delimiter)
		var actualFields []int
		for _, fieldId := range fieldIds {
			if fieldId-1 < len(inpLineSplit) {
				actualFields = append(actualFields, fieldId-1)
			}
		}

		for i, field := range actualFields {
			if i == len(actualFields)-1 {
				sb.Write([]byte(inpLineSplit[field]))
				break
			}
			sb.Write([]byte(inpLineSplit[field] + args.delimiter))
		}
		sb.Write([]byte("\n"))
	}
	return sb.String(), nil
}

func getFieldIds(fieldsString string) ([]int, error) {
	var fieldIds []int
	fields := strings.Split(fieldsString, ",")
	for _, field := range fields {
		splitField := strings.Split(field, "-")
		switch {
		case len(splitField) == 1:
			if i, err := strconv.Atoi(splitField[0]); err != nil {
				return nil, err
			} else {
				fieldIds = append(fieldIds, i)
			}
		case len(splitField) == 2:
			from, err := strconv.Atoi(splitField[0])
			if err != nil {
				return nil, errors.New("wrong field format")
			}

			to, err := strconv.Atoi(splitField[1])
			if err != nil {
				return nil, errors.New("wrong field format")
			}

			for i := from; i <= to; i++ {
				fieldIds = append(fieldIds, i)
			}
		default:
			return nil, errors.New("wrong field format")
		}
	}
	sort.Ints(fieldIds)
	return fieldIds, nil
}
