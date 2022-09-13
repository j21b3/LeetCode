package main

func Reverse(input []int) {
	i, j := 0, len(input)-1
	for i < j {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
	return
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue, ret, reverse := make([][]*TreeNode, 0), make([][]int, 0), false
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

		if reverse {
			Reverse(tmpr)
		}

		ret = append(ret, tmpr)
		if len(tmpq) > 0 {
			queue = append(queue, tmpq)
		}
		reverse = !reverse
	}
	return ret
}
