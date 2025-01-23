package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			i--
		}
	}

	return prefix
}

func main() {
	test1 := []string{"flower", "flow", "flight"}
	test2 := []string{"c", "acc", "ccc"}

	fmt.Println(longestCommonPrefix(test1))
	fmt.Println(longestCommonPrefix(test2))
}
