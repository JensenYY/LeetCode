package main

import (
	"fmt"
)

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	for _, v := range nums {
		index := abs(v) - 1
		nums[index] = -abs(nums[index])
	}
	ret := make([]int, 0, 0)
	for k, v := range nums {
		if v > 0 {
			ret = append(ret, k+1)
		}
	}
	return ret
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(nums)
}
