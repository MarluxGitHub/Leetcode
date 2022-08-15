/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
    var stack []int

	var pu = 0
	var po = 0

	var res = true

	for {
		if(len(stack) == 0 && pu == len(pushed) && po == len(popped)) {
			break
		}

		if(len(stack) != 0 && popped[po] == stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			po++
			continue
		}

		if(pu == len(pushed) && po != len(popped)) {
			res = false
			break
		}

		stack = append(stack, pushed[pu])
		pu++
	}

	return res
}
// @lc code=end

