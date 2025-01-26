package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0

	for _, value := range nums {
		result ^= value
	}

	return result
}

func main() {
	fmt.Println(1 ^ 2 ^ 1 ^ 3^ 3)
}
