package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	v := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println(reverse(v))
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	temp := head.Next
	temp.Next = head
	head.Next = reverse(nextHead)
	return temp
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		t := A[i]
		for j, z := 0, len(t)-1; j < z; j, z = j+1, z-1 {

			t[j], t[z] = rp(t[z]), rp(t[j])
		}
	}
	return A
}

func rp(x int) int {
	if x == 1 {
		return 0
	}
	return 1
}
