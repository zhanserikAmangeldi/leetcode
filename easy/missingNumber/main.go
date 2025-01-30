package main

import "sort"

func missingNumber(nums []int) int { // the most stupid solution
	res := 0
	max := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < max; i++ {
		if i == max {
			res = max
			break
		}
		if i != nums[i] {
			res = i
			break
		}
	}

	return res
}

func missingNumberBit(nums []int) int { // The best?
	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= i
	}

	for _, val := range nums {
		res ^= val
	}

	return res
}
