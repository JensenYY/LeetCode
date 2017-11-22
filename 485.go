package main

import ()

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	max := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] == 1 {
				count++
			}
			continue
		}
		if nums[i] == nums[i-1] && nums[i] == 1 {
			count++
		} else if nums[i] != nums[i-1] && nums[i] == 1 {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 0
		}
	}
	if max < count {
		max = count
	}
	return max
}
