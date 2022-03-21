/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */


// @lc code=start
func intervalIntersection(A [][]int, B [][]int) [][]int {
    ans := make([][]int,0)
    for k := range ans {
        ans[k] = make([]int,0)
    }
    i := 0
    j := 0
    for i < len(A) && j < len(B) {
        low := max(A[i][0],B[j][0])
        high := min(A[i][1],B[j][1])
        if low <= high {
            ans = append(ans,[]int {low,high})
        }
        if A[i][1] < B[j][1] {
            i++
        } else {
            j++
        }
    }
    return ans
}



func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}
// @lc code=end

