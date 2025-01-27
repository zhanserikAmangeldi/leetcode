package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	clearString := ""
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	for _, value := range s {
		if !(('A' <= value && value <= 'Z') || ('a' <= value && value <= 'z') || ('0' <= value && value <= '9'))  {
			clearString = strings.ReplaceAll(s, string(value), "")
		}
	}

    fmt.Println(clearString)
	cnt := len(clearString) - 1

	for index := 0; index < len(clearString)/2; index++ {
		if clearString[index] != clearString[cnt] {
			return false
		}
		cnt--
	}

	fmt.Println(s)

	return true
}