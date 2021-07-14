# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        elif l2 is None:
            return l1
        if l1.val<l2.val:
            p1=l1.next
            p2=l2
            head=l1
        else:
            p1=l1
            p2=l2.next
            head=l2
        cur=head
        while p1 is not None and p2 is not None:
            if p1.val<p2.val:
                cur.next=p1
                cur=cur.next
                p1=p1.next
            else:
                cur.next=p2
                cur=cur.next
                p2=p2.next
        if p1 is not None:
            cur.next=p1
        else:
            cur.next=p2
        return head