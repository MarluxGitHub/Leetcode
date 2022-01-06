package leetcode

/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	trustMap := make(map[int]int)
	for _, t := range trust {
		trustMap[t[0]]--
		trustMap[t[1]]++
	}

	for k, v := range trustMap {
		if v == n-1 {
			return k
		}
	}
	return -1
}
// @lc code=end

