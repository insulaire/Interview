package main

import (
	"fmt"
)

func main() {
	t := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	preorderTraversal(&t)
	fmt.Println()
	preorderTraversal2(&t)
	fmt.Println()
	preorderTraversal3(&t)

	fmt.Println(preorderTraversal4(&t))
	fmt.Println(preorderTraversal5(&t))
	fmt.Println(preorderTraversal6(&t))
	fmt.Println(preorderTraversal7(&t))
	fmt.Println(preorderTraversal8(&t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

//中序遍历
func preorderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	preorderTraversal2(root.Left)
	fmt.Println(root.Val)
	preorderTraversal2(root.Right)
}

//后序遍历
func preorderTraversal3(root *TreeNode) {
	if root == nil {
		return
	}
	preorderTraversal3(root.Left)
	preorderTraversal3(root.Right)
	fmt.Println(root.Val)
}

//前序遍历
func preorderTraversal4(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//中序遍历
func preorderTraversal5(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

//后序遍历
func preorderTraversal6(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}
	var last *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || last == node.Right {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			last = node
		} else {
			root = node.Right
		}
	}
	return result
}

//dfs
func preorderTraversal7(root *TreeNode) []int {
	res := []int{}

	dfs(root, &res)

	//res =divideAndConquer(root)
	return res
}
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func divideAndConquer(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

//bfs
func preorderTraversal8(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		list := []int{}
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)

	}
	return res
}
