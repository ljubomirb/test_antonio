package main

import (
	"strings"
)

//MostOccurences returns letter that occures the most times in a given string
func MostOccurences(testString string) string {

	var mostUsedLetter string
	var maxCounted int

	for _, s := range testString {

		singleLetter := string(s)

		singleLetterOccurences := strings.Count(testString, singleLetter)
		if singleLetterOccurences > maxCounted {
			maxCounted = singleLetterOccurences
			mostUsedLetter = singleLetter
		}

	}

	return mostUsedLetter
}
