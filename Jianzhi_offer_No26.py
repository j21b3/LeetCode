# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# 需要注意的是题目中A或B任一为空就为False
class Solution:
    def compare(self, A: TreeNode, B: TreeNode):
        if B is None:
            return True
        elif A is None and B is not None:
            return False
        if A.val != B.val:
            return False
        return self.compare(A.left, B.left) and self.compare(A.right, B.right)

    def isSubStructure(self, A: TreeNode, B: TreeNode) -> bool:
        if A is None or B is None:
            return False
        return self.compare(A, B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)
