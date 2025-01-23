package main

import "fmt"

func twoSumN2(nums []int, target int) []int { // TODO: Decrease BIG O NOTATION(Cur: N^2 -> Future: ?)
	result := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j}...)
				break
			}
		}
	}

	return result
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	numsMap := map[int]int{}

	for index, num := range nums {
		half := target - num
		_, ok := numsMap[half]
		if  ok {
			result = append(result, []int{index, numsMap[half]}...)
			break
		}
		numsMap[num] = index
	}

	return result
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}
	target1 := 9
	target2 := 6
	target3 := 6

	fmt.Println(twoSum(nums1, target1))
	fmt.Println(twoSum(nums2, target2))
	fmt.Println(twoSum(nums3, target3))
}
