# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None:
            return head
        before = None
        cur = head
        after = head.next
        while after is not None:
            cur.next = before
            before, cur = cur, after
            after = after.next
        cur.next = before
        return cur
