package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var recFunc func(A *TreeNode, B *TreeNode) bool
	recFunc = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return recFunc(A.Left, B.Left) && recFunc(A.Right, B.Right)
	}

	return (A != nil && B != nil) && (recFunc(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
