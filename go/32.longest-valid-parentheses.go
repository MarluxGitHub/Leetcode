/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
    	max := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
				if i-stack[len(stack)-1] > max {
					max = i - stack[len(stack)-1]
				}
			} else {
				stack = []int{i}
			}
		}
	}
	return max
}
// @lc code=end

