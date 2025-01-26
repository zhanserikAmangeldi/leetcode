package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [-10,-3,0,5,9]
func sortedArrayToBST(nums []int) *TreeNode {
	head := &TreeNode{
		Val: nums[0],
	}

	cur := head

	for _, value := range nums {
		if value > cur.Val {

			cur.Right = value
		}
	}

	return head
}

func main() {

}
