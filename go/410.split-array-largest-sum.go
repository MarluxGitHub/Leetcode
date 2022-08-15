/*
 * @lc app=leetcode id=410 lang=golang
 *
 * [410] Split Array Largest Sum
 */

// @lc code=start
func splitArray(nums []int, m int) int {

    max, sum := 0, 0
    for _, k := range nums {
        if max < k {
            max = k
        }
        sum += k
    }
    // the possible result is among: [max, sum]
    left, right := max, sum

    for left <= right {
        mid := (left + right) / 2
        // need how many arrarys to get no more than mid?
        t := match(nums, mid)

        // typical binary search condition
        if t > m{
            // move to right
            left = mid + 1
        }else{
            // move to left
            right = mid - 1
        }
    }
    return left
}

// find out: need how many arrys to match the target sum.
func match(nums []int, target int) int {
    n, sum :=1, 0

    for _, k := range nums {
        sum += k
        // if it's over than the target, (1)need one more array (2) reset sum
        if sum > target {
            sum = k
            n++
        }
    }
    return n
}
// @lc code=end

