/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
    l := len(nums)
    for i := 0; i < l; i++ {
        if nums[i] > 0 {
            if nums[nums[i]] < 0 {
                return nums[i]
            }
            nums[nums[i]] *= -1
        } else {
            if nums[-nums[i]] < 0 {
                return -nums[i]
            }
            nums[-nums[i]] *= -1
        }
    }
    return -1
}
// @lc code=end

