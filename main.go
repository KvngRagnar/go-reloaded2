package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	OutputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Println("error reading file", err)
		return
	}

	words := string(data)
	result := TextProcessor(words)

	err = os.WriteFile(OutputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Done Successfully")
}

func TextProcessor(words string) string {

	words = cases(words)
	words = convBase(words)
	words = fixArt(words)
	words = fixPunc(words)
	words = fixQuote(words)

	return words + "\n"
}
