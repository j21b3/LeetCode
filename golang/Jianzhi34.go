package main

func pathSum(root *TreeNode, target int) [][]int {
	ret, path := make([][]int, 0), make([]int, 0)
	var fun func(cur *TreeNode, tar int)
	fun = func(cur *TreeNode, tar int) {
		if cur == nil {
			return
		}
		path = append(path, cur.Val)
		i := tar - cur.Val
		if i == 0 && cur.Left == nil && cur.Right == nil {
			// 这里要注意深浅拷贝
			tmp := make([]int, len(path))
			copy(tmp, path)
			ret = append(ret, tmp)
		}
		fun(cur.Left, i)
		fun(cur.Right, i)
		path = path[:len(path)-1]
	}
	fun(root, target)
	return ret
}
