package main

import "strings"

func fixQuote(text string) string {

	words := strings.Split(text, "'")
	for i := 0; i < len(words); i++ {
		if i % 2 == 1 {
			words[i] = strings.TrimSpace(words[i])
		}
	}
	return strings.Join(words, "'")
}

