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
     return height(root) != -1
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := height(root.Left)
	rightCount := height(root.Right)

	if math.Abs(float64(leftCount - rightCount)) > 1 || leftCount == -1 || rightCount == -1{
		return - 1
	}

	return 1 + max(leftCount , rightCount)
	 
}
