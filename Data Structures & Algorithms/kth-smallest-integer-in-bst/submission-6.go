/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count := []int{}

	countK(root, k, &count)
	return count[k-1]
}

func countK(root *TreeNode, k int, res *[]int) {

	if root == nil {
		return 
	}
	if len(*res) == k {
		return
	}
	countK(root.Left, k, res)
	*res = append(*res, root.Val)
	countK(root.Right, k, res )

	

	return 
}
