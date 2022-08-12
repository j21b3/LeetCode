package main

func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	matrix := make([][]bool, m)

	for i := range matrix {
		matrix[i] = make([]bool, n)
	}

	matrix[0][0] = true

	// 当s取空字符串的时候，只有p的所有偶数位上的字符为*的时候才能完全匹配
	// 例如 "" 和 a*b*c*
	for j := 2; j < n; j = j + 2 {
		matrix[0][j] = matrix[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 由于.和字符都只表示一个字符，因此这俩一类，*单独一类
			if p[j-1] == '*' {
				//如果*取0次 那么直接看该行j-2的位置是不是匹配
				//如果*让之前的字符多出现一次能不能匹配
				//如果*让之前的.多出现一次能不能匹配
				if matrix[i][j-2] {
					matrix[i][j] = true
				} else if matrix[i-1][j] && s[i-1] == p[j-2] {
					matrix[i][j] = true
				} else if matrix[i-1][j] && p[j-2] == '.' {
					matrix[i][j] = true
				}
			} else {
				// 如果是个普通字符，则直接判断是不是一样
				// 如果是. 则看上一个位置是不是匹配上了
				if matrix[i-1][j-1] && s[i-1] == p[j-1] {
					matrix[i][j] = true
				} else if matrix[i-1][j-1] && p[j-1] == '.' {
					matrix[i][j] = true
				}
			}
		}
	}
	return matrix[m-1][n-1]
}
