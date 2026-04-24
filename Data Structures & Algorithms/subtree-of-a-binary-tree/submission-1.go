/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false 
	}
	if root.Val == subRoot.Val && campare(root,subRoot) {
		return true
	}
	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}

func campare(root, sub *TreeNode) bool {
	if sub == nil && root == nil {
		return true
	}
	if root != nil && sub != nil && root.Val == sub.Val {
		return campare(root.Left,sub.Left) && campare(root.Right,sub.Right) 
	}
	return false
}