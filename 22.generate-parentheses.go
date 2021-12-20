/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int,  s string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}

	if left > 0 {
		dfs(left-1, right, s+"(", res)
	}

	if right > left {
		dfs(left, right-1, s+")", res)
	}
}
// @lc code=end

