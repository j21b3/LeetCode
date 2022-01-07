# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        leftNode = rightNode = head
        for i in range(n):
            rightNode = rightNode.next
        if not rightNode:
            return head.next  # 这里别忘了要特判
        while rightNode.next:
            rightNode = rightNode.next
            leftNode = leftNode.next
        leftNode.next = leftNode.next.next
        return head