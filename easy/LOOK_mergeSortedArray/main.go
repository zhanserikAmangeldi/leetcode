package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Create a new slice to store merged result
	result := make([]int, m+n)
	
	// Initialize pointers for nums1, nums2 and result
	i, j, k := 0, 0, 0
	
	// Compare and merge while elements exist in both arrays
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements from nums1 if any
	for i < m {
		result[k] = nums1[i]
		i++
		k++
	}
	
	// Copy remaining elements from nums2 if any 
	for j < n {
		result[k] = nums2[j]
		j++
		k++
	}
	
	// Copy result back to nums1
	copy(nums1, result)
}

func main() {
	nums1 := []int{}
	nums2 := []int{1}

	merge(nums1, 0, nums2, 1)
}
