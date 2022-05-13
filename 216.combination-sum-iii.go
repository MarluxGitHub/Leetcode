/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
var ans [][]int
func combinationSum3(k int, n int) [][]int {
    ans = [][]int{}
    find([]int{}, []int{1,2,3,4,5,6,7,8,9}, n, k)
    return ans
}

func find(cur []int, nums []int, t int, k int) {
    if len(cur) == k {
        return
    }
    if len(nums) == 0 {
        return
    }

    if t <= 0 {
        return
    }
    for i, v := range nums {
        if v == t {
            temp := []int{}
            temp = append(temp, cur...)
            temp = append(temp, v)
            if len(temp) == k {
                ans = append(ans, temp)
                return
            }
            continue
        }
        temp := []int{}
        temp = append(temp, nums[i+1:]...)
        find(append(cur, v), temp, t-v, k)
    }
}
// @lc code=end

