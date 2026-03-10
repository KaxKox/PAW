package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}