package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
null/nil 异常处理
dummy node 哑巴节点
快慢指针
插入一个节点到排序链表
从一个链表中移除一个节点
翻转链表
合并两个链表
找到链表的中间节点
*/
func main() {
	v := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	fmt.Println(add(v, 3))
}

//向有序列表中插入x
func add(head *ListNode, x int) *ListNode {
	node := &ListNode{Val: x}
	if head == nil {
		return node
	}
	dummy := &ListNode{Val: math.MaxInt64}
	dummy.Next = head
	head = dummy
	//查找插入节点
	for head.Next != nil && head.Next.Val < x {
		head = head.Next
	}
	//保存下一个元素
	temp := head.Next
	//插入节点连接下一个元素
	node.Next = temp
	//插入节点连接
	head.Next = node

	return dummy.Next
}

//删除节点为X的节点
func delete(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for head.Next != nil && head.Next.Val == x {
		//删除该节点
		head.Next = head.Next.Next
	}
	return dummy.Next
}

//翻转链表
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		//缓存下个元素
		temp := head.Next
		//断链并指向上一个元素
		head.Next = prev
		//保存当前元素
		prev = head
		//移到下个元素继续处理
		head = temp
	}
	return prev
}

//判断链表是否有环 有则返回环开始节点
func hasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// fast如果初始化为head.Next则中点在slow.Next
	// fast初始化为head,则中点在slow
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			//相遇后slow 继续向前走
			//fast改为起点同步运行
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
