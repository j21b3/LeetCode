# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        ret = []
        q1=[root]
        q2=[]
        while q1:
            tmp=[]
            q2=q1[:]
            q1.clear()
            while q2:
                tmp.append(q2[0].val)
                if q2[0].left is not None:
                    q1.append(q2[0].left)
                if q2[0].right is not None:
                    q1.append(q2[0].right)
                del q2[0]
            ret.append(tmp)
        return ret