package distance

import (
	"regexp"
	"strings"
)

// LetterPairs Get all of the pairs of letters for a string
func LetterPairs(str string) []string {
	numPairs := len(str) - 1
	pairs := make([]string, numPairs)

	for i := 0; i < numPairs; i++ {
		pairs[i] = str[i : i+2]
	}
	return pairs
}

// WordLetterPairs Get all of the pairs in all of the words for a string
func WordLetterPairs(str string) []string {
	allPairs := make([]string, 0)
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(str, -1)

	for i := 0; i < len(words); i++ {
		pairs := LetterPairs(words[i])
		for _, pair := range pairs {
			allPairs = append(allPairs, pair)
		}
	}
	return allPairs
}

// Sanitize Perform some sanitization steps
func Sanitize(str string) string {
	return strings.Replace(strings.ToLower(str), " ", "", -1)
}

// Compare two strings, and spit out a number from 0-1
func Compare(str1, str2 string) float32 {
	sanitizedStr1 := Sanitize(str1)
	sanitizedStr2 := Sanitize(str2)
	pairs1 := WordLetterPairs(sanitizedStr1)
	pairs2 := WordLetterPairs(sanitizedStr2)
	intersection := 0
	union := len(pairs1) + len(pairs2)
	if union == 0 {
		if sanitizedStr1 == sanitizedStr2 {
			return 1
		}
		return 0
	}

	for i := 0; i < len(pairs1); i++ {
		pair1 := pairs1[i]
		for j := 0; j < len(pairs2); j++ {
			pair2 := pairs2[j]
			if strsEquals := strings.Compare(pair1, pair2); strsEquals == 0 {
				intersection++
				pairs2 = append(pairs2[:j], pairs2[j+1:]...)
				break
			}
		}
	}
	return 2 * float32(intersection) / float32(union)
}
