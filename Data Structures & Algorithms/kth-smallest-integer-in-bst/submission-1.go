/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import ("slices")
func kthSmallest(root *TreeNode, k int) int {
    count := []int{}

	countK(root, k, &count)
	slices.Sort(count)
	return count[k-1]
}

func countK(root *TreeNode, k int, res *[]int) {

	if root == nil {
		return 
	}

	countK(root.Left, k, res)
	
	countK(root.Right, k, res )

	*res = append(*res, root.Val)

	return 
}
