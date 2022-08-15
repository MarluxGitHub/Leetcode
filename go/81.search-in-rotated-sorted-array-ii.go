/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {

	left := 0
	right := len(nums) - 1
	last := -999

	for left < len(nums) {
		if (nums[left] == target) {
			return true
		}

		if last <= nums[left] {
			last = nums[left]
		} else {
			break;
		}


		left++
	}

	for right > left {
		if nums[right] == target {
			return true
		}

		if last >= nums[right] {
			last = nums[right]
		} else {
			break;
		}


		right--
	}

	return false
}
// @lc code=end

