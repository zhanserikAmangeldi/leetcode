package main

import "fmt"

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	p := 3
	pp := 2
	cur := 5

	for i := 4; i < n; i++ {
		pp = p
		p = cur
		cur = p + pp
	}

	return cur 

	// [0][1][1][2][3][5][8][13]
	//       [1][2][3][4][5][6]

}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(7))
	fmt.Println(climbStairs(8))
	fmt.Println(climbStairs(9))
}
