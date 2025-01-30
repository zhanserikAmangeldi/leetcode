package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	inMind := 0
	result := []byte{}
	indexA := len(a) - 1
	indexB := len(b) - 1

	for indexA >= 0 || indexB >= 0 || inMind == 1 {
		if indexA >= 0 {
			inMind += int(a[indexA] - '0')
			indexA--
		}

		if indexB >= 0 {
			inMind += int(b[indexB] - '0')
			indexB--
		}

		result = append([]byte((strconv.Itoa(inMind % 2))), result...)
		inMind = inMind / 2
	}

	return string(result)

}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("10", "1"))
	fmt.Println(addBinary("11", "1"))
}
