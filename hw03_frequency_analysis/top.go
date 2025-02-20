package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	mapWordsCount := make(map[string]int)
	wordsInText := strings.Fields(text)

	for _, word := range wordsInText {
		mapWordsCount[word]++
	}

	wordKeys := make([]string, 0, len(mapWordsCount))
	for key := range mapWordsCount {
		wordKeys = append(wordKeys, key)
	}

	sort.Slice(wordKeys, func(i, j int) bool {
		if mapWordsCount[wordKeys[i]] == mapWordsCount[wordKeys[j]] {
			return wordKeys[i] < wordKeys[j]
		}
		return mapWordsCount[wordKeys[i]] > mapWordsCount[wordKeys[j]]
	})

	needLen := 10
	if needLen > len(wordKeys) {
		needLen = len(wordKeys)
	}

	result := wordKeys[:needLen]
	return result
}
