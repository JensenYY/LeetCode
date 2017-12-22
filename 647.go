package main

import (
	"fmt"
	"strings"
)

func countSubstrings(s string) int {
	if strings.Compare(s, "") == 0 {
		return 0
	}
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			count++
		}
		for j := 0; i-1-j >= 0 && i+j < n && s[i-1-j] == s[i+j]; j++ {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countSubstrings("aa"))
}
