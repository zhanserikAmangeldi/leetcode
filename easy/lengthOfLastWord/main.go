package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	result := 0

	for _, value := range s {
		if value != 32 {
			result++
		} else {
			result = 0
			continue
		}
	}

	return result
}

func main() {

	fmt.Println(' ')
	fmt.Println(lengthOfLastWord("H f"))
	fmt.Println(lengthOfLastWord(" fly  me    to  the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy   "))
}
