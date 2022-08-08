class Solution:
    def longestValidParentheses(self, s) -> int:
        stack = []
        res = 0
        for i in range(len(s)):
            if not stack or s[i] == '(' or s[stack[-1]] == ')':
                stack.append(i)
            else:
                stack.pop()
                res = max(res, i - (stack[-1] if stack else - 1))
        return res



if __name__ == "__main__":
    st = "()(()"
    lu = Solution()
    t = lu.longestValidParentheses(st)
    print(t)
