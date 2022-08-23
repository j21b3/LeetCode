package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur, next := head, head
	cur = nil
	for next != nil {
		t := next.Next
		next.Next = cur
		cur = next
		next = t
	}
	return cur

}
