# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mirrorTree(self, root: TreeNode) -> TreeNode:
        def mirr(head: TreeNode):
            if head is None:
                return
            elif head.left is None and head.right is None:
                return
            head.left, head.right = head.right, head.left
            mirr(head.left)
            mirr(head.right)

        mirr(root)
        return root
