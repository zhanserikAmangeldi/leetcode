package main

func containsDuplicate(nums []int) bool {
	vals := map[int]int{}

	for _, value := range nums {
		if vals[value] > 0 {
			return true
		}

		vals[value] += 1
	}

	return false
}
