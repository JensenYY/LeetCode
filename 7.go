package main

import (
	"fmt"
)

func reverse(x int) int {
	ans := 0
	isNega := false
	if x < 0 {
		isNega = true
		x = 0 - x
	}
	for i := x; i > 0; i /= 10 {
		ans = ans*10 + (i % 10)
	}
	if x > 2147483647 {
		return 0
	}
	if isNega {
		return -ans
	}
	return ans
}

func main() {
	fmt.Println(reverse(123))
}
