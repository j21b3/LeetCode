package main

func exist(board [][]byte, word string) bool {
	steps := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(board), len(board[0])

	//当前位置（i，j），字符串需要匹配第k位
	var dfs func(i, j, cur int) bool
	dfs = func(i, j, cur int) bool {
		if board[i][j] != word[cur] {
			return false
		}
		if cur == len(word)-1 {
			return true
		}

		board[i][j] = '.'
		for _, s := range steps {
			nexti, nextj := i+s[0], j+s[1]

			//若nexti，nextj超出范围，或者下一个位置已经去过了
			if nexti < 0 || nexti >= m || nextj < 0 || nextj >= n || board[nexti][nextj] == byte('.') {
				continue
			}
			if res := dfs(nexti, nextj, cur+1); res {
				return true
			}

		}
		board[i][j] = word[cur]
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res := dfs(i, j, 0); res {
				return true
			}
		}
	}
	return false
}
