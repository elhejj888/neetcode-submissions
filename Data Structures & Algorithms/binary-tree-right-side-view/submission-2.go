/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
		max := queue[0].Val
		for _ = range(len(queue)){
			curr := queue[0]
			queue = queue[1:]
			
			if curr.Val > max {
				max = curr.Val
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		}
		vars = append(vars, max)
	}
	return vars
	
}
