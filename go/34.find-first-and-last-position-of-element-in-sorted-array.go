/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return []int{findLeft(nums, target, left, mid), findRight(nums, target, mid, right)}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func findLeft(nums []int, target int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func findRight(nums []int, target int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// @lc code=end

