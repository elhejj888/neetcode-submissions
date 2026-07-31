/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	leftCount := 1
	rightCount := 1
	if root == nil {
		return 0
	} 

		
		rightCount += maxDepth(root.Right)

		leftCount += maxDepth(root.Left)

	return max(leftCount, rightCount)
    
}
