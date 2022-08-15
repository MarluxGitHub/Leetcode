/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 { return -1 }

    currStepQueue := make([][]int, 0)
    nextStepQueue := make([][]int, 0)
    step := 0
    target := len(grid) - 1

    currStepQueue = append(currStepQueue, []int{0, 0})
    grid[0][0] = 1

    for len(currStepQueue) > 0 {
        step++
        for _, cell := range currStepQueue {
            if cell[0] == target && cell[1] == target { return step }

            for n := -1; n < 2; n++ {
                for m := -1; m < 2; m++ {
                    nt := cell[0] + n
                    mt := cell[1] + m
                    if nt < 0 || nt > target || mt < 0 || mt > target || grid[nt][mt] == 1 { continue }
                    grid[nt][mt] = 1
                    nextStepQueue = append(nextStepQueue, []int{nt, mt})
                }
            }
        }
        currStepQueue, nextStepQueue = nextStepQueue, currStepQueue[:0]
    }

    return -1
}
// @lc code=end

