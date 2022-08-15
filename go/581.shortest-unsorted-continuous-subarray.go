/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
    if len(nums)==1 {
        return 0
    }
    dup:=make([]int,len(nums))
    copy(dup, nums)
    sort.Ints(dup)
    start, end := 0, len(nums)-1
    for start<end {
        if nums[start]==dup[start] {
            start++
        }
        if nums[end]==dup[end] {
            end--
        }
        if nums[start]!=dup[start]&&nums[end]!=dup[end] {
            break
        }
    }
    if start==end {
        return 0
    }
    return end-start+1
}
// @lc code=end

