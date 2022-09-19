package main

func permutation(s string) []string {
	ret := make([]string, 0)
	m := make(map[byte]int)
	ele := ""

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			ret = append(ret, ele)
			return
		}
		for k, _ := range m {
			if m[k] != 0 {
				ele += string(k)
				m[k]--
				dfs(idx + 1)
				ele = ele[:len(ele)-1]
				m[k]++
			}
		}
	}

	dfs(0)
	return ret
}
