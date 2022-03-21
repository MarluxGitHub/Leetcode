/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
    	if len(nums) < 2 {
		return 0
	}
	var result int
	current := 0
	next := 0
	for i := 0; i < len(nums)-1; i++ {
		current = max(current, nums[i]+i)
		if i == next {
			next = current
			result++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

