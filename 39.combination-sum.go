/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
    	res := [][]int{}
	dfs(candidates, target, 0, []int{}, &res)
	return res
}

func dfs(candidates []int, target, index int, path []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := index; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs(candidates, target - candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}
// @lc code=end

