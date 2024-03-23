package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			result.Next = list1
			list1 = list1.Next
		} else {
			result.Next = list2
			list2 = list2.Next
		}
		result = result.Next
	}

	if list1 == nil && list2 != nil {
		result.Next = list2
	}

	if list2 == nil && list1 != nil {
		result.Next = list1
	}

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	r := mergeTwoLists2(l1, l2)
	fmt.Println(r)
}
