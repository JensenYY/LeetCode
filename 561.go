package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i += 2 {
		count += nums[i]
	}
	return count
}