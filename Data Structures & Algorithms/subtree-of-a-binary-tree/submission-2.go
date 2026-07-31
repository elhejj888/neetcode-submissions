/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil{
		return true
	}
	if root == nil || subRoot == nil{
		return false
	}
	isTrue := false
    if root.Val == subRoot.Val {
		isTrue = isEqual(root, subRoot)
	}
	if isTrue {
		return isTrue
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p==nil || q==nil || p.Val != q.Val {
		return false
	} 
	return isEqual(p.Left, q.Left) && isEqual(p.Right , q.Right)
}
