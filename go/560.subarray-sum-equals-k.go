package leetcode

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
    dict := map[int]int{}
    count := 0
    curr := 0
    dict[0] = 1
    for _, each := range nums {
        curr = curr + each
        count = count + dict[curr-k]
        dict[curr] = dict[curr] + 1
    }
    return count
}
// @lc code=end

