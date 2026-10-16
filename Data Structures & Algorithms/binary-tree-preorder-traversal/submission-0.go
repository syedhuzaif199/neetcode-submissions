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
    out = append(out, root.Val)
    out = recur(root.Left, out)
    out = recur(root.Right, out)
    return out
}

func preorderTraversal(root *TreeNode) []int {
    out := []int{}
    return recur(root, out)
}
