/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(node *TreeNode, low *int, high *int) bool {
    if node == nil {
        return true
    }

    if low != nil && node.Val <= *low {
        return false
    }

    if high != nil && node.Val >= *high {
        return false
    }

    return validate(node.Left, low, &node.Val) &&
           validate(node.Right, &node.Val, high)
}
