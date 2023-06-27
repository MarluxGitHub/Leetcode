package leetcode

import "sort"

/*
 * @lc app=leetcode id=1502 lang=golang
 *
 * [1502] Can Make Arithmetic Progression From Sequence
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
   // sort array
   sort.Ints(arr)

   for i := 2; i < len(arr); i++ {
		if arr[i] - arr[i-1] != arr[i-1] - arr[i-2] {
			return false
		}
	}

	return true

}
// @lc code=end

