package main

//给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路 递归生成每一种可能性 然后拼接

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := generate(start, i-1)
		right := generate(i+1, end)
		//生成树顶为i的所有子树
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				root := &TreeNode{Val: i}
				root.Left = left[j]
				root.Right = right[k]
				res = append(res, root)
			}
		}
	}
	return res
}
