package main

import (
	"fmt"
)

/*
	XOR
	1. 0 ^ N = N
	2. N ^ N = 0
*/
func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 2, 1}
	fmt.Println(singleNumber(arr))
}
