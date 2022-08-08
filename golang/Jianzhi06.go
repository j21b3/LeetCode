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

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	ret := []int{}
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	reverse(ret)
	return ret
}

func reverse(lis []int) {
	i, j := 0, len(lis)-1
	for i < j {
		lis[i], lis[j] = lis[j], lis[i]
		i++
		j--
	}
}
