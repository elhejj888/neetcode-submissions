/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	
	matrix := [][]int{}
	count := 0
	if root == nil {
		return matrix
	}
	
	insert(root, &matrix, count)
	return matrix

}

func insert(root *TreeNode, matrix *[][]int, count int){
	if root == nil {
		return
	}
	if len(*matrix) <= count {
		(*matrix) = append(*matrix, [] int {root.Val})
		
	}else {
	(*matrix)[count] = append((*matrix)[count], root.Val)
	}
	localCount := count + 1
	insert(root.Left, matrix, localCount)
	insert(root.Right, matrix, localCount)
	

	return
}


