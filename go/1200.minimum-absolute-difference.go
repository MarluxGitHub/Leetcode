package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=1200 lang=golang
 *
 * [1200] Minimum Absolute Difference
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

	min := 2147483647

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] - arr[i] < min {
			min = arr[i+1] - arr[i]
		}
	}

	res := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}
// @lc code=end

