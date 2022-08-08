class Solution:
    def __init__(self):
        self.ret = []

    def trace(self, candidates, target: int, result, p: int):
        if target == 0:
            ret = result[:]
            self.ret.append(ret)
            return
        for i in range(p, len(candidates)):
            if candidates[i] <= target:
                result.append(candidates[i])
                self.trace(candidates, target - candidates[i], result, i)
                result.pop()
            else:
                break

    def combinationSum(self, candidates, target: int):
        candidates.sort()

        resu = []
        self.trace(candidates, target, resu, 0)
        return self.ret


S = Solution()
candidates = [2, 3, 6, 7]
target = 7
print(S.combinationSum(candidates, target))
