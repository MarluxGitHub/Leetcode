 /*
 * @lc app=leetcode id=994 lang=kotlin
 *
 * [994] Rotting Oranges
 */

package lt994

import java.util.LinkedList/*
import java.util.Queue

// @lc code=start
class Solution {
    fun orangesRotting(grid: Array<IntArray>): Int {
        // step 1
        if (grid.isEmpty()) return 0
        val queue: Queue<Pair<Int, Int>> = LinkedList()
        var fresh = 0
        var minutes = 0
        // step 2
        for (i in grid.indices) {
            for (j in grid[i].indices) {
                when (grid[i][j]) {
                    0 -> {}
                    1 -> fresh++
                    else -> queue.offer(i to j)
                }
            }
        }

        // result
        if (fresh == 0) return 0 // no fresh
        else if (queue.isEmpty()) return -1 // no rotten
        else if (queue.size == grid.size * grid[0].size) return 0 // all are rotten

        // step 3
        while (queue.isNotEmpty()) {
            var size = queue.size
            while (size != 0) {
                val (i, j) = queue.poll()
                if (isValid(grid, i - 1, j)) { // left
                    grid[i - 1][j] = 2
                    queue.offer(i - 1 to j)
                    --fresh
                }
                if (isValid(grid, i + 1, j)) { // right
                    grid[i + 1][j] = 2
                    queue.offer(i + 1 to j)
                    --fresh
                }
                if (isValid(grid, i, j - 1)) { // top
                    grid[i][j - 1] = 2
                    queue.offer(i to j - 1)
                    --fresh
                }
                if (isValid(grid, i, j + 1)) { // bottom
                    grid[i][j + 1] = 2
                    queue.offer(i to j + 1)
                    --fresh
                }
                --size
            }
            minutes++
        }
        // result
        return if (fresh == 0) minutes - 1 else -1
    }
}

private fun isValid(grid: Array<IntArray>, i: Int, j: Int) = i in grid.indices && j in grid[i].indices && grid[i][j] == 1

// @lc code=end

