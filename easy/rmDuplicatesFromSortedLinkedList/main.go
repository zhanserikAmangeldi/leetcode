package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		if cur.Next != nil {
			if cur.Val == cur.Next.Val {
				temp := cur.Next.Next
				cur.Next = temp
			} else {
				cur = cur.Next
			}
		}
	}

	return head
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	// list2 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 3,
	// 		Next: &ListNode{
	// 			Val:  4,
	// 			Next: nil,
	// 		},
	// 	},
	// }

	resList := deleteDuplicates(list1)

	fmt.Println(resList)
	for resList != nil {
		fmt.Println(resList.Val)
		resList = resList.Next
	}
}
