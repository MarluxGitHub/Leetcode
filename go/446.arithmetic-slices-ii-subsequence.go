package leetcode

/*
 * @lc app=leetcode id=446 lang=golang
 *
 * [446] Arithmetic Slices II - Subsequence
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
    	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += dp[j][diff] + 1
			res += dp[j][diff]
		}
	}
	return res
}
// @lc code=end

