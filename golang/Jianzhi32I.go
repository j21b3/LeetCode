package main

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue, ret := make([]*TreeNode, 0, 0), make([]int, 0, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ret = append(ret, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return ret
}
