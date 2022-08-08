# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        ret = []
        q=[root]
        while q:
            ret.append(q[0].val)
            if q[0].left is not None:
                q.append(q[0].left)
            if q[0].right is not None:
                q.append(q[0].right)
            del q[0]
        return ret