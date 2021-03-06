package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		count := m[word]
		m[word] = count + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

