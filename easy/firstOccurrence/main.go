package main

import "fmt"

func strStr(haystack string, needle string) int {

	if needle == haystack {
		return 0
	}
	if len(needle) == 0 || len(needle) > len(haystack) {
		return -1
	}

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		flag := true

		for j := 0; j < lenNeedle; j++ {
			fmt.Println(i, j)
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}

	return -1
}

func main() {
	haystack1 := "mississippi"
	needle1 := "pi"

	// haystack2 := "leetcode"
	// needle2 := "leeto"

	fmt.Println(strStr(haystack1, needle1))
	// fmt.Println(strStr(haystack2, needle2))
}
