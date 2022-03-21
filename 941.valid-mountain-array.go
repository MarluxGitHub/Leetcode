package leetcode

/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 */

// @lc code=start
func validMountainArray(arr []int) bool {
    inscreasing, descreasing := false, false

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			inscreasing = true
			if descreasing {
				return false
			}
		} else if arr[i] < arr[i-1] {
			descreasing = true
		} else {
			return false
		}
	}

	return inscreasing && descreasing
}
// @lc code=end

