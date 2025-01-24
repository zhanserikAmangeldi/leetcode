package main

import "fmt"

func isValid(s string) bool {
	arr := []rune{}

	par := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, value := range s {

		if char, ok := par[value]; ok {
			if len(arr) > 0 && arr[len(arr)-1] == char {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		} else {
			arr = append(arr, value)
		}
	}

	// for i := 0; i < len(s); i++ {
	// 	if par[top] == s[i] && len(arr) != 0 {
	// 		arr = arr[:len(arr)-1]
	// 	} else {
	// 		arr = append(arr, s[i])
	// 	}
	// 	if len(arr) != 0 {
	// 		top = arr[len(arr)-1]
	// 	}
	// }

	return len(arr) == 0
}

func main() {
	test1 := "(){}}{}"
	test2 := "()[]{}" // = "(", "()"
	test3 := "]"
	test4 := "([])"

	fmt.Println(isValid(test1))
	fmt.Println(isValid(test2))
	fmt.Println(isValid(test3))
	fmt.Println(isValid(test4))
}
