package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}
	return head.Next
}

func main() {
	l1 := new(ListNode)
	l11 := new(ListNode)
	l12 := new(ListNode)

	l1.Val = 1
	l11.Val = 3
	l12.Val = 6

	l1.Next = l11
	l11.Next = l12

	l2 := new(ListNode)
	l21 := new(ListNode)
	l22 := new(ListNode)

	l2.Val = 1
	l21.Val = 2
	l22.Val = 5

	l2.Next = l21
	l21.Next = l22

	l3 := mergeTwoLists(l1, l2)

	for nil != l3 {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}
