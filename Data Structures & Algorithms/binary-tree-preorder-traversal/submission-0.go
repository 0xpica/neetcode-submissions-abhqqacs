/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}
	arr = append(arr,root.Val)
	l := preorderTraversal(root.Left)
	r := preorderTraversal(root.Right)
	arr = append(arr,append(l,r...)...)
	return arr
}
