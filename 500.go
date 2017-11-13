package main

import (
	"strings"
)

const (
	s1 = "qwertyuiopQWERTYUIOP"
	s2 = "asdfghjklASDFGHJKL"
	s3 = "zxcvbnmZXCVBNM"
)

func findWords(words []string) []string {
	var newWords []string = make([]string, 0, len(words))

	for _, word := range words {
		var tmp string = strings.ToLower(word)
		var c1, c2, c3 int = 0, 0, 0
		var length int = len(word)

		for _, c := range tmp {
			if strings.ContainsRune(s1, c) {
				c1++
			} else if strings.ContainsRune(s2, c) {
				c2++
			} else if strings.ContainsRune(s3, c) {
				c3++
			}
		}
		if c1 == length || c2 == length || c3 == length {
			newWords = append(newWords, word)
		}
	}
	return newWords
}
