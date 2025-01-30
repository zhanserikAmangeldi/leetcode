package main

func findDisappearedNumbers(nums []int) []int {
	m := map[int]int{}
	res := []int{}

	for index, value := range nums {

	}

	for key, value := range m {
		if value == 0 {
			res = append(res, key)
		}
	}

	return res

	
}

