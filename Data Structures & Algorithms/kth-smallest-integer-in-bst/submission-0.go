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
	if root == nil {
		return 0
	}
	countK(root, k, &count)
	fmt.Println(count)
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
