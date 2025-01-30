package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	cur := head
	slow := head
	fast := head
	mid := 0

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		mid++
	}

	slow = Reverse(slow)

	for cur != nil && slow != nil {
		if cur != slow {
			return false
		}
		cur = cur.Next
		slow = slow.Next
	}

	return true
}
