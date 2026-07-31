/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	getGoodNodes(root, root.Val, &res)
	return res
}
func getGoodNodes(node *TreeNode, parentVal int, res *int) {
	if node == nil {
		return
	}
	if node.Val >= parentVal{
		*res++
		parentVal = node.Val
	}
	getGoodNodes(node.Left, parentVal, res)
	getGoodNodes(node.Right, parentVal, res)
	return
}
