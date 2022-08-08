class Solution:
    def __init__(self):
        self.stack = []
        self.ret = []
        self.flag = 0

    def recursion(self, n, left_remain, right_remain):
        # case 1 : add ( if possible
        if n == 0 and left_remain == 0 and right_remain == 0:
            self.ret.append("".join(self.stack))
            return
        if left_remain > 0:
            self.stack.append('(')
            self.flag += 1
            self.recursion(n, left_remain - 1, right_remain)
            self.stack.pop()
            self.flag -= 1
        if right_remain > 0 and self.flag > 0:
            self.stack.append(')')
            self.flag -= 1
            self.recursion(n - 1, left_remain, right_remain - 1)
            self.stack.pop()
            self.flag += 1

    def generateParenthesis(self, n: int):
        self.recursion(n, n, n)
        return set(self.ret)


s = Solution()
print(s.generateParenthesis(3))
