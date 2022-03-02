/*
 * @lc app=leetcode id=39 lang=javascript
 *
 * [39] Combination Sum
 */

// @lc code=start
/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
    var res = []

    dfs(candidates, target, 0, [], res)

    return res
};

function dfs(candidates, target, index, path, res) {
    if(target < 0) return

    if (target === 0) {
        res.push(path.slice())
        return
    }
    for (var i = index; i < candidates.length; i++) {
        path.push(candidates[i])
        dfs(candidates, target - candidates[i], i, path, res)
        path.pop()
    }
}
// @lc code=end

