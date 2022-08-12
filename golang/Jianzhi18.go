package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Next: head,
	}
	before, cur := dummy, head
	for cur != nil && cur.Val != val {
		cur = cur.Next
		before = before.Next
	}

	before.Next = cur.Next
	cur.Next = nil

	return dummy.Next
}
