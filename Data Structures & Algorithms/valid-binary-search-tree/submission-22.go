/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}
func validate(root *TreeNode, minAllowed *int, maxAllowed *int) bool {
	if root == nil{
		return true
	}
	if minAllowed != nil && root.Val <= *minAllowed{
		return false
	}
	if maxAllowed != nil && root.Val >= *maxAllowed{
		return false
	}

	return validate(root.Left, minAllowed, &root.Val) && validate(root.Right, &root.Val, maxAllowed)
}