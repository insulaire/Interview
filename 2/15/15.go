package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
func main() {
	node := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	fmt.Println(TreeToDoublyList(&node)) //输出为 1 2 3 4 5 的双向链表

	fmt.Println(levelOrder(&node))
}

var head *TreeNode //头部节点
var pre *TreeNode  //上次迭代的节点

func TreeToDoublyList(root *TreeNode) *TreeNode {
	//
	dfs(root)
	//尾部处理
	pre.Right = head
	//头部处理
	head.Left = pre
	return head
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	//中序遍历 先处理左边节点
	dfs(root.Left)

	if pre == nil {
		//未迭代过 确定头部节点
		head = root
	} else {
		//上次处理的节点的right节点指向当前节点
		pre.Right = root
	}
	//当前节点指向上次处理的节点
	root.Left = pre
	//保存当前处理的节点
	pre = root
	//中序遍历 后处理右边节点
	dfs(root.Right)
}

//、、请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}
	flag := true
	for len(stack) > 0 {
		l := len(stack)
		arr := make([]int, len(stack))
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if flag {
				arr[i] = node.Val
			} else {
				arr[l-i-1] = node.Val
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		flag = !flag
		ans = append(ans, arr)
	}
	return ans
}
