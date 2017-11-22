package main

import ()

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	last := 0
	cur := 0
	for cur < len(nums) {
		if nums[cur] != 0 {
			nums[last], nums[cur] = nums[cur], nums[last]
			last++
		}
		cur++
	}
}
