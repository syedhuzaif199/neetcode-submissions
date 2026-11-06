/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
func f(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	hl, bl := f(root.Left)
	hr, br := f(root.Right)
	h := 1 + max(hl, hr)
	b := bl && br && abs(hl-hr) <= 1
	return h, b
}

func isBalanced(root *TreeNode) bool {
	_, b := f(root)
	return b
}
