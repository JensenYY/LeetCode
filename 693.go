package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	flag := -1
	var mod int
	for i := n; i > 0; i /= 2 {
		mod = i % 2
		if mod != flag {
			flag = mod
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(7))
}
