package main

import "fmt"

// func romanToInt(s string) int {
// 	sum := 0
// 	romans := map[string]int{
// 		"I":  1,
// 		"IV": 4,
// 		"V":  5,
// 		"IX": 9,
// 		"X":  10,
// 		"XL": 40,
// 		"L":  50,
// 		"XC": 90,
// 		"C":  100,
// 		"CD": 400,
// 		"D":  500,
// 		"CM": 900,
// 		"M":  1000,
// 	}

// 	for i := 0; i < len(s)-1; i++ {
// 		if romans[string(s[i])] < romans[string(s[i+1])] {
// 			sum -= romans[string(s[i])]
// 		} else {
// 			sum += romans[string(s[i])]
// 		}
// 	}

// 	sum += romans[string(s[len(s)-1])]

// 	return sum
// }

func romanToInt(s string) int {
	sum := 0
	prev := 0
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// MCMXCIV =  0 VICXMCM
	for i := len(s) - 1; 0 <= i; i-- {
		cur := romans[s[i]]

		if prev > cur {
			sum -= cur
		} else {
			sum += cur
		}

		prev = cur
	}

	return sum
}

func main() {
	firstTest := "III"
	secondTest := "LVIII" // IIIVL
	thirdTest := "MCMXCIV"

	fmt.Println(romanToInt(firstTest))
	fmt.Println(romanToInt(secondTest))
	fmt.Println(romanToInt(thirdTest))
}
