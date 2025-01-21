package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
)

func Top10(inputStr string) []string {
	words := getWords(inputStr)
	wordsCount := getWordsCount(words)
	wordsList := sortWordsByCount(wordsCount)
	if len(wordsList) > 10 {
		return wordsList[:10]
	}
	return wordsList
}

func getWords(inputStr string) []string {
	words := strings.Fields(inputStr)
	return words
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
		return cmp.Compare(wordsCount[a], wordsCount[b]) * -1 // sorting inversion
	}
	wordsList := make([]string, 0, len(wordsCount))
	for k := range wordsCount {
		wordsList = append(wordsList, k)
	}
	slices.Sort(wordsList)
	slices.SortStableFunc(wordsList, countSort)
	return wordsList
}
