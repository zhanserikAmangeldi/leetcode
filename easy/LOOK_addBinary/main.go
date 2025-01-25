package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		b, a = a, b
	}

	cntB := len(b) - 1
	temp := 0

	for i := len(a) - 1; i >= 0; i-- {
		
	}

	// 11010
	//   101

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("10", "1"))
	fmt.Println(addBinary("11", "1"))
}
