/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recur(root *TreeNode, out []int) []int {
    if root == nil {
        return out
    }
    out = recur(root.Left, out)
    out = recur(root.Right, out)
    out = append(out, root.Val)
    return out
}

func postorderTraversal(root *TreeNode) []int {
    out := []int{}
    return recur(root, out)
}
