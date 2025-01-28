package main

import "strconv"

func isHappy(n int) bool {
	for {
		if n == 1 || n == 7 {
			return true
		} else if (n >= 2 && n < 7) || (n > 7 && n <= 9) {
			return false
		}

		strN := strconv.Itoa(n)
		n = 0
		for i := 0; i < len(strN); i++ {
			digit, _ := strconv.Atoi(string(strN[i]))
			n += digit * digit
		}
	}
}
