package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	str := ""
	cur := head
	res := 0
	for cur != nil {
		str += strconv.Itoa(cur.Val)
		cur = cur.Next
	}

	for i, val := range str {
		res = res + (int(val-'0') << (len(str) - i - 1))
	}

	return res
}
