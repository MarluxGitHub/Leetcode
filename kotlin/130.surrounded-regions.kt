/*
 * @lc app=leetcode id=130 lang=kotlin
 *
 * [130] Surrounded Regions
 */

// @lc code=start
class Solution {
    fun solve(board: Array<CharArray>): Array<CharArray> {

        if(board.size < 3 || board.get(0).size < 3) return board

        for(y in board.indices) {
            if(board.get(y).get(0) == 'O') dfs(board, y, 0)
            if(board.get(y).get(board.get(0).size-1) == 'O') dfs(board, y, board.get(0).size-1)
        }

        for(x in board.get(0).indices) {
            if(board.get(0).get(x) == 'O') dfs(board, 0, x)
            if(board.get(board.size-1).get(x) == 'O') dfs(board, board.size-1, x)
        }

        for(y in board.indices) {
            for(x in board.get(0).indices) {
                if(board.get(y).get(x) == 'A') board[y][x] = 'O'
                else if(board.get(y).get(x) == 'O') board[y][x] = 'X'
            }
        }

        return board
    }

    fun dfs(board: Array<CharArray>, y:Int, x:Int) {
        if(!(y in board.indices) || !(x in board.get(y).indices) || board.get(y).get(x) != 'O') return
        board[y][x] = 'A'
        dfs(board, y+1, x)
        dfs(board, y-1, x)
        dfs(board, y, x+1)
        dfs(board, y, x-1)
    }
}
// @lc code=end

