class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(self, l1, l2):
        dummy = ListNode(0)  # 构造虚节点
        move = dummy  # 设置移动节点等于虚节点
        while l1 and l2:  # 都不空时
            if l1.val < l2.val:
                move.next = l1  # 移动节点指向数小的链表
                l1 = l1.next
            else:
                move.next = l2
                l2 = l2.next
            move = move.next
        move.next = l1 if l1 else l2  # 连接后续非空链表
        return dummy.next  # 虚节点仍在开头

    def mergeKLists(self, lists) -> ListNode:
        head = None
        for each in lists:
            head = self.mergeTwoLists(head, each)
        return head


t = [[1, 4, 5], [1, 3, 4], [2, 6]]
s = Solution()
print(s.mergeKLists(t))
