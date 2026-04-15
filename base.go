package main

import(

	"strconv"
	"strings"
)

func convBase(word string) string {

	words := strings.Fields(word)
	for b := 0; b < len(words); b++ {
	
		if words[b] == "(bin)" && b > 0 {

			num, err := strconv.ParseInt(words[b-1], 2, 64)
			if err == nil {

				words[b-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:b], words[b+1:]...)
		} 
		if words[b] == "(hex)" && b > 0{
			num, err := strconv.ParseInt(words[b-1], 16, 64)
			if err == nil {
				words[b-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:b], words[b+1:]...)
		}
	}
	return strings.Join(words, " ")

}