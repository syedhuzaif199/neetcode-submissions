/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func f(root *TreeNode) (int, int) {
    if root == nil {
        return  -1, -1
    }
	hl, dl := f(root.Left)
	hr, dr := f(root.Right)
	h := 1 + max(hl, hr)
	d := 2 + hl + hr
	return h, max(d, max(dl, dr))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, d := f(root)
	return d
}