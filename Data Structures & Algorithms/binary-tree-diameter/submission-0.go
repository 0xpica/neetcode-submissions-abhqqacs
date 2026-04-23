/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxHeight(root.Left)
	r := maxHeight(root.Right)
	d := l + r 

	sub := max(diameterOfBinaryTree(root.Left),diameterOfBinaryTree(root.Right))

	return max(d,sub)
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0 
	}
	return 1 + max(maxHeight(root.Left),maxHeight(root.Right))
}
