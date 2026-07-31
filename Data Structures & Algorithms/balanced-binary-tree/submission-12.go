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
	fmt.Println("right : ", right, " left : ", left)
	if right - left > 1 || right - left < -1 {
		return false
	}


	return isBalanced(root.Left) && isBalanced(root.Right)
	
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := height(root.Left)
	rightCount := height(root.Right)

	return 1 + max(leftCount , rightCount)
	 
}
