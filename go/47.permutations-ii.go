package leetcode

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    return backtrack(nums, make([]bool, len(nums)), []int{}, [][]int{})
}

func backtrack(nums []int, used []bool, permutation []int, result [][]int) [][]int {
    if len(permutation) == len(nums) {
        return append(result, append([]int{}, permutation...))
    }
    for i := 0; i < len(nums); i++ {
        if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        used[i] = true
        permutation = append(permutation, nums[i])
        result = backtrack(nums, used, permutation, result)
        permutation = permutation[:len(permutation)-1]
        used[i] = false
    }
    return result
}
// @lc code=end

