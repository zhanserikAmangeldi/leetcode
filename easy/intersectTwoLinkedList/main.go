package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	intersection := map[*ListNode]int{}
	i := 0
	for headA != nil {
		i++
		intersection[headA] = i
        headA = headA.Next
	}

	for headB != nil {
		if intersection[headB] > 0 {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
