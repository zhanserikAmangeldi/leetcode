package main

func majorityElement(nums []int) int {
	mapNums := map[int]int{}

	for _, num := range nums {
		mapNums[num]++
		if mapNums[num] > len(nums)/2 {
			return num
		}
	}

	return 0
}
