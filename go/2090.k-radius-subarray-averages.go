/*
 * @lc app=leetcode id=2090 lang=golang
 *
 * [2090] K Radius Subarray Averages
 */

// @lc code=start
func getAverages(nums []int, k int) []int {
    res := make([]int, len(nums))
	if len(nums) < 2*k {
		for i := 0; i < len(nums); i++ {
			res[i] = -1
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if i < k || i > len(nums)-k-1 {
				res[i] = -1
			} else {
				sum := 0
				for j := i - k; j < i+k+1; j++ {
					sum += nums[j]
				}
				res[i] = sum / (2*k + 1)
			}
		}
	}

	return res
}
// @lc code=end

