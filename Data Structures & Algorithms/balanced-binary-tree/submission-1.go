/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if abs(l,r) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0 
	}
	return 1 + max(getHeight(root.Left),getHeight(root.Right))
}

func abs(n1,n2 int) int {
	if n1 > n2 {
		return n1-n2
	}
	return n2-n1
}