package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
)

func Top10(inputStr string) []string {
	words := strings.Fields(inputStr)
	wordsCount := getWordsCount(words)
	wordsList := sortWordsByCount(wordsCount)
	if len(wordsList) > 10 {
		return wordsList[:9]
	}
	return wordsList
}

func getWordsCount(wordsList []string) map[string]int {
	wordsCount := make(map[string]int)
	for _, word := range wordsList {
		wordsCount[word] += 1
	}
	return wordsCount
}

func sortWordsByCount(wordsCount map[string]int) []string {
	countSort := func(a, b string) int {
		if order := cmp.Compare(wordsCount[a], wordsCount[b]); order == 0 {
			return order
		} else {
			return order
		}
	}
	wordsList := make([]string, 0, len(wordsCount))
	for k := range wordsCount {
		wordsList = append(wordsList, k)
	}
	slices.SortFunc(wordsList, countSort)
	slices.Reverse(wordsList)
	return wordsList
}
