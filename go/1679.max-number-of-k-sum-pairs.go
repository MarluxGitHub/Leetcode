package leetcode

import "sort"

/*
 * @lc app=leetcode id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) int {
    if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	l,r := 0, len(nums)-1
	erg := 0

	for l < r {
		if nums[l]+nums[r] == k {
			erg++
			l++
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			r--
		}
	}

	return erg
}
// @lc code=end

