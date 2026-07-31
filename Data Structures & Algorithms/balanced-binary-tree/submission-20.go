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
    right , left := height(root.Right), height(root.Left)
	if math.Abs(float64(right - left)) > 1 || left == -1 || right == -1 {
		return false
	}
	
	return true
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
