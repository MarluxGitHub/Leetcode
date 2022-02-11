/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
    	var row, col, block [9]map[byte]bool
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		block[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := row[i][board[i][j]]; ok {
				return false
			}
			row[i][board[i][j]] = true
			if _, ok := col[j][board[i][j]]; ok {
				return false
			}
			col[j][board[i][j]] = true
			if _, ok := block[i/3*3+j/3][board[i][j]]; ok {
				return false
			}
			block[i/3*3+j/3][board[i][j]] = true
		}
	}
	return true
}
// @lc code=end

