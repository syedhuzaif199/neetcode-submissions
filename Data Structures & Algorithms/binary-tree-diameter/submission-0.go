/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max_depth(root *TreeNode) int {
    if root == nil {
        return  -1
    }
    return 1 + max(max_depth(root.Left), max_depth(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return -1
    }
    dia := max_depth(root.Left) + max_depth(root.Right) + 2
    dia2 := diameterOfBinaryTree(root.Left)
    dia3 := diameterOfBinaryTree(root.Right)
    return max(dia, max(dia2, dia3))
}
