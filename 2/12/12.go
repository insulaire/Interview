package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	//边界
	if root == nil {
		return root
	}

	if root.Val > key {
		//向左遍历
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		//向右遍历
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//删除此节点
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}

	return root
}
