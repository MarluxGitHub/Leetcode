/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
    if(len(nums) == 0) {
		return 0
	}

	max := nums[0]
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		cur = maxInt(nums[i], cur + nums[i])
		max = maxInt(max, cur)
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

