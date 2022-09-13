package main

/*
	由于输入postorder是个后续遍历，要找的原树是一个二叉搜索树，所以可以有如下结论：
		左节点<=根节点<=右节点
		postorder倒着读是一个反着的前序遍历：根节点 右孩子 左孩子
	所以对一个postorder，从最后一位开始倒着读，首先是根节点，然后是右子树，然后是左子树
	其右子树中的所有元素都比根节点大，左子树都比其小
	所以当遍历碰到一个小的时候，那个元素一定是数值最相近的节点左孩子
	所以拿一个单调栈存变大的元素，遇到小的之后就把它对应的父亲找出来，直到这个栈为空
*/

func verifyPostorder(postorder []int) bool {
	stack, root := make([]int, 0), 0xffffffff
	for j := len(postorder) - 1; j > -1; j-- {
		if postorder[j] > root {
			return false
		}
		for len(stack) != 0 && postorder[j] < stack[len(stack)-1] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[j])
	}
	return true
}
