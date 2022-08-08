class Solution:

    def can_attack(self, x, y):
        if (x + y) in self.diag1 or (x - y) in self.diag2 or y in self.row:
            return 1
        return 0

    def gen_graph(self, n):
        ret = []
        for i in range(n):
            ret.append('.' * n)
        for x, y in self.queen_position:
            tmp = list(ret[x])
            tmp[y] = 'Q'
            ret[x] = ''.join(tmp)
        return ret

    def trace(self, x, n):
        if x == n - 1:
            for y in range(n):
                if self.can_attack(x, y):
                    continue
                self.queen_position.append((x, y))
                self.answer.append(self.gen_graph(n))
                self.queen_position.pop()
            return
        for y in range(n):
            if self.can_attack(x, y):
                continue
            self.diag1.append(x + y)
            self.diag2.append(x - y)
            self.row.append(y)
            self.queen_position.append((x, y))
            self.trace(x + 1, n)
            self.diag1.pop()
            self.diag2.pop()
            self.row.pop()
            self.queen_position.pop()

    def solveNQueens(self, n: int):
        self.diag1 = []  # 斜向右上(x+y)
        self.diag2 = []  # 斜向右下(x-y)
        self.row = []  # 纵向(y)
        self.queen_position = []
        self.answer = []
        self.trace(0, n)
        return self.answer


if __name__ == "__main__":
    C = Solution()
    print(C.solveNQueens(4))
