/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
	queue := []*TreeNode{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {

        res = append(res, queue[0].Val)

        levelSize := len(queue)

        for i := 0; i < levelSize; i++ {
            curr := queue[0]
            queue = queue[1:]

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }
        }
    }
	return res
	
}
