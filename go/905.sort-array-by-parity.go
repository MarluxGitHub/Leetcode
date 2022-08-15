package leetcode

/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {

    start, end :=0 , len(nums) - 1

    for start <end{
        if nums[start]%2 > nums[end]%2 {
			nums[start], nums[end] = nums[end], nums[start]
        }

        if nums[start]%2==0 {start+=1}
        if nums[end]%2==1 {end-=1}
    }
    return nums
}
// @lc code=end

