# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        a = head
        b = head.next
        if a.val==val:
            return b
        while b is not None and b.val!=val:
            a,b=b,b.next
        if b is None:
            return head
        a.next=b.next
        del b
        return head