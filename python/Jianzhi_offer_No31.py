class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        stack = []
        i, j = 0, 0
        while i < len(pushed):
            if not stack and pushed[i] == popped[j]:
                i += 1
                j += 1
            elif stack and stack[-1] != popped[j]:
                stack.append(pushed[i])
                i += 1
            elif stack and stack[-1] == popped[j]:
                del stack[-1]
                j += 1
            else:
                stack.append(pushed[i])
                i += 1
        while j < len(pushed) and stack:
            if stack and stack[-1] == popped[j]:
                del stack[-1]
                j += 1
            else:
                return False
        return False if stack else True


