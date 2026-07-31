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
	isFalse := true
    right , left := height(root.Right, &isFalse), height(root.Left, &isFalse)
	fmt.Println("right : ", right, " left : ", left)
	if right - left > 1 || right - left < -1 {
		return false
	}
	
	return isFalse
	
}
func height(root *TreeNode, isFalse *bool) int {
	if root == nil {
		return 0
	}
	leftCount := height(root.Left, isFalse)
	rightCount := height(root.Right, isFalse)

	if leftCount - rightCount > 1 || leftCount - rightCount < -1{
		*isFalse = false
	}

	return 1 + max(leftCount , rightCount)
	 
}
