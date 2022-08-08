# 剑指 Offer 12. 矩阵中的路径

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def dfs(i, j, k):   # 设置内置函数可以不通过传参直接使用上层函数的变量
            if not 0 <= i < len(board) or not 0 <= j < len(board[0]) or board[i][j] != word[k]:
                return False
            if k == len(word) - 1:
                return True
            board[i][j] = ''  # 防止回走将匹配上的字符修改
            res = dfs(i + 1, j, k + 1) or dfs(i - 1, j, k + 1) or dfs(i, j + 1, k + 1) or dfs(i, j - 1, k + 1)
            board[i][j] = word[k]
            return res

        for i in range(len(board)):
            for j in range(len(board[0])):
                if dfs(i,j,0) : return True
        return False