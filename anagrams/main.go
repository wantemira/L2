package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Printf("%v", Anagrams(words))
}

func Anagrams(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		key := sortWord(word)
		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string)

	for _, group := range groups {
		if len(group) < 2 {
			continue
		}

		sort.Strings(group)
		result[group[0]] = group
	}

	return result
}

func sortWord(word string) string {
	runes := []rune(word)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
