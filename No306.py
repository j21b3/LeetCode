class Solution:
    def __init__(self):
        self.numlist = []

    def trace(self, p, num):
        if p >= len(num):
            return True
        plu = self.numlist[-1] + self.numlist[-2]
        for i in range(p + 1, len(num) + 1):
            if num[p] == '0' and i - p > 1:
                return False
            t = int(num[p:i])
            if t > plu:
                break
            elif t == plu:
                self.numlist.append(t)
                r = self.trace(i, num)
                if r:
                    return True
                self.numlist.pop()
        return False

    def isAdditiveNumber(self, num: str) -> bool:
        if len(num) < 3:
            return False
        for num0 in range(1, len(num)):
            if num[0] == '0' and num0 > 1:
                break
            self.numlist.append(int(num[0:num0]))
            for num1 in range(num0 + 1, len(num)):
                if num[num0] == '0' and num1 - num0 > 1:
                    break
                self.numlist.append(int(num[num0:num1]))
                ret = self.trace(num1, num)
                if ret:
                    return True
                self.numlist.pop()
            self.numlist.pop()
        return False


num = "0235813"
a = Solution()
print(a.isAdditiveNumber(num))
