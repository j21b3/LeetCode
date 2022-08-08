# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        p1 = head
        p2 = head
        flag = 0
        while p2 is not None and flag < k:
            p2 = p2.next
            flag += 1
        if flag < k and p2 is None:
            return None
        while p2 is not None:
            p1 = p1.next
            p2 = p2.next
        return p1
