/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import ("slices")
func rightSideView(root *TreeNode) []int {
    vars := []int{}
	queue := []*TreeNode{}
	if root == nil {
		return vars
	}
	queue = append(queue, root)
	for {
		if len(queue) == 0{
			break
		}
		tmp := [] int {}
		for _ = range(len(queue)){
			curr := queue[0]
			queue = queue[1:]
			tmp = append(tmp, curr.Val)
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		}
		fmt.Println(queue, "----", tmp)
		vars = append(vars, slices.Max(tmp))
	}
	return vars
	
}
