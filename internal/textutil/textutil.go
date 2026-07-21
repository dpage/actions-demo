// Package textutil provides small string helpers used by the demo CLI.
package textutil

import (
	"strings"
	"unicode"
)

// Reverse returns s with its runes in reverse order, so multi-byte characters
// survive the trip intact.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards, having
// first discarded anything that is not a letter or a digit and folded case.
func IsPalindrome(s string) bool {
	cleaned := Normalise(s)
	return cleaned == Reverse(cleaned)
}

// Normalise lowercases s and strips every rune that is neither a letter nor a
// digit.
func Normalise(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// WordCount returns the number of whitespace-separated words in s.
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// Title returns s with the first letter of each whitespace-separated word
// upper-cased and the remainder left alone.
func Title(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
