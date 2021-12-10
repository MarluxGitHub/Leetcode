/*
 * @lc app=leetcode id=1306 lang=golang
 *
 * [1306] Jump Game III
 */

// @lc code=start
func canReach(arr []int, start int) bool {
	visited := make(map[int]bool)
	return dfs(arr, start, visited)
}

func dfs(arr []int, start int, visited map[int]bool) bool {
	if start < 0 || start >= len(arr) || visited[start] {
		return false
	}
	if arr[start] == 0 {
		return true
	}
	visited[start] = true
	return dfs(arr, start+arr[start], visited) || dfs(arr, start-arr[start], visited)
}
// @lc code=end

