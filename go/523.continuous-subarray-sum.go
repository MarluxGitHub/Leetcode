package leetcode

/*
 * @lc app=leetcode id=523 lang=golang
 *
 * [523] Continuous Subarray Sum
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
    sum := 0
    m := make(map[int]int)

    for i, num := range nums {
        sum = (sum + num) % k

		if(sum == 0 && i > 0) {
			return true
		}

        if pre, ok := m[sum]; ok {
            if i-pre > 1 {
                return true
            }
        } else {
            m[sum] = i
        }
    }
    return false
}
// @lc code=end

