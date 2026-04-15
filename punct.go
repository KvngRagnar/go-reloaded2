package main

import "regexp"

func fixPunc(text string) string {
	p1:= regexp.MustCompile(`\s+([?.,;:!]+)`)
	p2:= regexp.MustCompile(`([?.,:;!]+)\s`)

	text = p1.ReplaceAllString(text, "$1 ")
	text = p2.ReplaceAllString(text, "$1")

	return text
}