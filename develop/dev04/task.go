package main

import (
	"sort"
	"strings"
)

func GetAnagrams(words *[]string) *map[string][]string {
	allSortedWords := make(map[string]string)
	for _, word := range *words {
		lowerWord := strings.ToLower(word)
		wordSorted := sortString(lowerWord)
		found := false
		for _, v := range allSortedWords {
			if v == wordSorted {
				found = true
				break
			}
		}
		if !found {
			allSortedWords[lowerWord] = wordSorted
		}
	}

	result := make(map[string][]string, len(allSortedWords))
	for _, word := range *words {
		lowerWord := strings.ToLower(word)
		wordSorted := sortString(lowerWord)
		for k, v := range allSortedWords {
			if wordSorted == v {
				result[k] = append(result[k], lowerWord)
			}
		}
	}

	for k, v := range result {
		if len(v) == 1 {
			delete(result, k)
		} else {
			sort.Strings(result[k])
		}
	}

	return &result
}

func sortString(word string) string {
	wordRunes := []rune(word)
	sort.Slice(wordRunes, func(i, j int) bool {
		return wordRunes[i] < wordRunes[j]
	})

	wordSorted := string(wordRunes)
	return wordSorted
}
