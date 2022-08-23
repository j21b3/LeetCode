package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head := ListNode{}
	cur := &head
	cl1, cl2 := l1, l2
	for cl1 != nil && cl2 != nil {
		if cl2.Val < cl1.Val {
			cur.Next = cl2
			cl2 = cl2.Next
		} else {
			cur.Next = cl1
			cl1 = cl1.Next
		}
		cur = cur.Next
	}
	if cl1 != nil {
		cur.Next = cl1
	} else {
		cur.Next = cl2
	}
	return head.Next
}
