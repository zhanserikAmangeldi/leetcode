package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	cur := head
	nodeMap := make(map[*ListNode]int)

	for cur != nil {
		nodeMap[cur] += 1
		if nodeMap[cur] > 1 {
			return true
		}
	}

	return false
}

func main() {

}
