/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
    left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left <= right {

		switch {
			case nums[mid] == target:
				return mid
			case nums[mid] > target:
				right = mid - 1
			default:
				left = mid + 1
		}
		mid = (left + right) / 2
	}

	if nums[mid] > target {
		return mid
	}
	return mid + 1
}
// @lc code=end

