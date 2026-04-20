/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}
	l := postorderTraversal(root.Left)
	r := postorderTraversal(root.Right)
	arr = append(arr,append(l,r...)...)
	arr = append(arr,root.Val)
	return arr
}
