package main

import "strings"

func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		counts[word]++
	}

	return counts
}
