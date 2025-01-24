package main

import "fmt"

func removeElement(nums []int, val int) int {
	result := len(nums) - 1
	t := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val && result >= i {
			temp := nums[result]
			nums[result] = nums[i]
			nums[i] = temp
			i--
			result--
			t++
		}
	}

	return len(nums) - t
}

func main() {
	test1 := []int{3, 2, 2, 3}
	test2 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(test1, 3))
	fmt.Println(removeElement(test2, 2))
}
