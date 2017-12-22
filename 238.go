package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums), len(nums))
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= right
		right *= nums[i]
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
