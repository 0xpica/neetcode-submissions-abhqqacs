/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}	


	arr := []int{}
	l := inorderTraversal(root.Left)
	arr = append(arr, l...)
	arr = append(arr, root.Val)
	r := inorderTraversal(root.Right)
	arr = append(arr, r...)
	return arr
}
