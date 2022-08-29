package leetcode

/*
 * @lc app=leetcode id=659 lang=golang
 *
 * [659] Split Array into Consecutive Subsequences
 */

// @lc code=start
type Chain struct {
    start int
    end int
}

func isPossible(nums []int) bool {
    var chains []Chain
    a: for _, n := range nums {
        if len(chains) == 0 {
            chains = append(chains, Chain{start: n, end: n})
            continue
        }
        for k := len(chains)-1; k >= 0; k-- {
            if chains[k].end + 1 == n {
                chains[k].end = n
                continue a
            }
        }
        chains = append(chains, Chain{start: n, end: n})
    }
    for _, c := range chains {
        if c.end - c.start < 2 {
            return false
        }
    }
    return true
}
// @lc code=end

