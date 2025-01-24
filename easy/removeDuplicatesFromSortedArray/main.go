package main

import "fmt"

func removeDuplicates(nums []int) int {
	result := 1

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums, i)
		if nums[i-1] != nums[i] {
			nums[result] = nums[i]
			result++
		}
	}

	return result
}

func main() {
	test1 := []int{1, 2}
	test2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(test1))
	fmt.Println(removeDuplicates(test2))
}
