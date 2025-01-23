package main

import (
	"fmt"
	"strconv"
)

// func isPalindrome(x int) bool {
// 	s := strconv.Itoa(x)
// 	count1 := 0
// 	count2 := len(s) - 1

// 	for i := 0; i < len(s)/2+1; i++ {
// 		if s[count1] != s[count2] {
// 			return false
// 		}
// 		count1++
// 		count2--
// 	}

// 	return true
// }

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2+1:]
	if len(s)%2 == 0 {
		secondHalf = s[len(s)/2:]
	}
	reversedSecondHalf := ""

	for i := len(secondHalf) - 1; i >= 0; i-- {
		reversedSecondHalf += string(secondHalf[i])
	}

	fmt.Println(firstHalf)
	fmt.Println(secondHalf)

	return firstHalf == reversedSecondHalf

}

func main() {
	firstTest := -1211
	secondTest := -121
	thirdTest := 10

	fmt.Println(isPalindrome(firstTest))
	fmt.Println(isPalindrome(secondTest))
	fmt.Println(isPalindrome(thirdTest))

}
