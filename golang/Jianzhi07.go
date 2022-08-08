package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
* 前序遍历的第一个节点一定是根节点，构建左子树流程和构建一个树一样
* 根据中序遍历中根节点的位置分两边，左边的是左子树的中序遍历，右子树同理
* 根据前序遍历的性质，先构建根节点，然后递归构建左子树、右子树即可
 */

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val: preorder[0],
	}

	i := 0
	for ; i < len(inorder); i++ { //搜索中序遍历
		if inorder[i] == node.Val {
			break
		}
	}

	node.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return node
}
