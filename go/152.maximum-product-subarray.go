/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mx, mn := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn = max(nums[i], max(nums[i]*mx, nums[i]*mn)), min(nums[i], min(nums[i]*mx, nums[i]*mn))
		res = max(res, mx)
	}

	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
// @lc code=end

