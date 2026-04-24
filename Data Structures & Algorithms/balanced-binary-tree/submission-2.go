/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	balanced, _ := dfs(root)
	return balanced
}

func dfs(root *TreeNode) (bool,int) {
	if root == nil {
		return true , 0  
	}
	lb, lh := dfs(root.Left)
	rb, rh := dfs(root.Right) 
	balanced := lb && rb && abs(lh,rh) <= 1
	height := 1 + max(lh,rh)
	return balanced, height
}

func abs(n1,n2 int) int {
	if n1 > n2 {
		return n1-n2
	}
	return n2-n1
}