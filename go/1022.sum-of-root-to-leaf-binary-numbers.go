package leetcode

/*
 * @lc app=leetcode id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumRootToLeaf(root *TreeNode) int {
	s := 0
	if root != nil {
		depth(root, 0*2+root.Val, &s)
	}
	return s
}

func depth(r *TreeNode, weight int, sum *int) {
	if r != nil {
		rl, rr := r.Left, r.Right
		if rl == nil && rr == nil {
			*sum = *sum + weight
		}
		if rl != nil {
			depth(rl, weight*2+rl.Val, sum)
		}
		if rr != nil {
			depth(rr, weight*2+rr.Val, sum)
		}
	}
}

// @lc code=end

