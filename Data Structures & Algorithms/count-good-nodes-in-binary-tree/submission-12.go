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
	getGoodNodes(root, root.Val ,root.Val, &res)
	return res
}
func getGoodNodes(node *TreeNode, rootVal int, parentVal int, res *int) {
	if node == nil {
		return
	}
	if node.Val >= rootVal && node.Val >= parentVal{
		*res++
		parentVal = node.Val
	}
	getGoodNodes(node.Left, rootVal, parentVal, res)
	getGoodNodes(node.Right, rootVal, parentVal, res)
	return
}
