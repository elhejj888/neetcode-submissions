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

    left := height(root.Left)
    if left == -1 {
        return -1
    }

    right := height(root.Right)
    if right == -1 {
        return -1
    }

    if math.Abs(float64(left-right)) > 1 {
        return -1
    }

    return 1 + max(left, right)
}
