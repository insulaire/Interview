package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return valid(root.Left, lower, root.Val) && valid(root.Right, root.Val, upper)
}
