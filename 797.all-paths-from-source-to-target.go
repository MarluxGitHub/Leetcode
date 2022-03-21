/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
		res := make([][]int, 0)
		cur := make([]int, 0)
		cur = append(cur, 0)
		dfs(graph, 0, cur, &res)
		return res
}

func dfs(graph [][]int, cur int, curPath []int, res *[][]int) {
		if cur == len(graph) - 1 {
				tmp := make([]int, len(curPath))
				copy(tmp, curPath)
				*res = append(*res, tmp)
				return
		}
		for _, next := range graph[cur] {
				tmp := make([]int, len(curPath))
				copy(tmp, curPath)
				tmp = append(tmp, next)
				dfs(graph, next, tmp, res)
		}
}
// @lc code=end

