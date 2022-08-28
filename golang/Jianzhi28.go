package main

func isSymmetric(root *TreeNode) bool {
	var judge func(l, r *TreeNode) bool
	judge = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		} else if l != nil && r != nil {
			return l.Val == r.Val && judge(l.Left, r.Right) && judge(l.Right, r.Left)
		} else {
			return false
		}
	}
	if root == nil {
		return true
	}
	return judge(root.Left, root.Right)
}
