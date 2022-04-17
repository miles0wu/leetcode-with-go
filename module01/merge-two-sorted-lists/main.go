package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// list1, list2 都空才停
	head := &ListNode{0, nil}
	protect := head
	for list1 != nil && list2 != nil {
		var nextNode *ListNode
		if list1.Val <= list2.Val {
			nextNode = list1
			list1 = list1.Next
		} else {
			nextNode = list2
			list2 = list2.Next
		}

		head.Next = nextNode
		head = head.Next
	}

	if list1 == nil {
		head.Next = list2
	} else {
		head.Next = list1
	}

	return protect.Next
}
