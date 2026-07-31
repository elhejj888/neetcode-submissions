/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

    if root.Val == subRoot.Val && isEqual(root, subRoot) {
    return true
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
