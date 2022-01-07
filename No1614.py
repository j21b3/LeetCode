class Solution:
    def maxDepth(self, s: str) -> int:
        p, size = 0, 0
        for word in s:
            if word == '(':
                p += 1
                size = max(size, p)
            elif word == ')':
                p -= 1
        return size
