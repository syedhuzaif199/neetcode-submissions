/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    return sameTree(root, subRoot) ||
        sameTree(root.Left, subRoot)  ||
        sameTree(root.Right, subRoot) ||
        isSubtree(root.Left, subRoot) ||
        isSubtree(root.Right, subRoot)
}
