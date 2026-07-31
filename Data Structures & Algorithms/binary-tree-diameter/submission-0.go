/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	getHeight(root, &diameter)
	return diameter
}

func getHeight(root *TreeNode, diameter *int) int {

	if root == nil {
		return 0
	} 

	rightCount := getHeight(root.Right, diameter)

	leftCount := getHeight(root.Left, diameter)

	*diameter = max(*diameter, rightCount + leftCount)

	return 1 + max(leftCount,rightCount)
}



