package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue, ret := make([][]*TreeNode, 0), make([][]int, 0)
	queue = append(queue, []*TreeNode{root})
	for len(queue) > 0 {
		tmpq, tmpr := make([]*TreeNode, 0), make([]int, 0)
		lay := queue[0]
		queue = queue[1:]

		for _, each := range lay {
			tmpr = append(tmpr, each.Val)
			if each.Left != nil {
				tmpq = append(tmpq, each.Left)
			}
			if each.Right != nil {
				tmpq = append(tmpq, each.Right)
			}
		}
		ret = append(ret, tmpr)
		if len(tmpq) > 0 {
			queue = append(queue, tmpq)
		}
	}
	return ret
}
