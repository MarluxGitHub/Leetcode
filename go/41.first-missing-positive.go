/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
	num := 1;
	current := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || current == nums[i] {
			continue
		}

		if nums[i] == num {
			current = nums[i]
			num++
		} else {
			return num
		}
	}
	return num
}
// @lc code=end

