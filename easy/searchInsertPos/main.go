package main

import "fmt"

func searchInsert(nums []int, target int) int {
		start, end := 0, len(nums)-1
		mid := ((start + end) / 2) + 1

		if nums[start] > target {
			return 0
		}

		if nums[end] < target {
			return len(nums)
		}

		for end > start {
			mid = ((start + end) / 2) + 1
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				start = mid + 1
			} else if nums[mid] > target {
				end = mid - 1
			}
		}

		if nums[end] < target {
			return end + 1
		}

		return end

}

func main() {
	fmt.Println("First test: ")
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

	fmt.Println("Second test: ")
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 1))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 3))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 6))

	fmt.Println("Third test: ")
	fmt.Println(searchInsert([]int{2, 3, 5, 6, 9}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 3))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 1))
}
